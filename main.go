package main

import (
	"Dandelion/config"
	db "Dandelion/db/service"
	"Dandelion/handler"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conf, err := config.LoadConfig("config/.")
	if err != nil {
		log.Fatalf("Load config file failed: %v", err)
	}

	dbConnect, err := sqlx.Connect(conf.Postgres.Driver, conf.Postgres.DatabaseSource())
	if err != nil {
		log.Fatalf("Connect database err: %v\n", err)
	}
	defer dbConnect.Close()

	rdb := initRedisClient(conf.RedisClient.Addr)
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Connect Redis error: %v", err)
	}

	queries := db.NewQueries(dbConnect)

	router := handler.NewHandler(conf, queries, rdb)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Router,
	}

	runDBMigration(conf.Postgres.MigrationUrl, conf.Postgres.DatabaseUrl())

	shutdown(ctx, srv)

}

func shutdown(ctx context.Context, srv *http.Server) {
	// 启动HTTP服务器
	go func() {
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("Error starting server: ", err)
		}
	}()

	// 用于接收退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	/// 等待退出信号
	<-quit
	log.Println("Received exit signal, shutting down...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server shutdown: ", err)
	}

	log.Println("Server closed...")
}

// sql file migrate
func runDBMigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("cannot create new migrate instance, Err: %v", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrate up, Err: %v", err)
	}

	log.Print("db migrated successfully...")
}

func initRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

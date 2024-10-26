package db

import (
	"Dandelion/db/model"
	"context"
	"log"
)

func (q *Queries) AddFriendRecord(ctx context.Context, userId, friendId uint32) error {

	sql := `INSERT INTO friends (user_id, friend_id) VALUES ($1, $2)`

	return q.db.QueryRowContext(ctx, sql, userId, friendId).Err()

}

func (q *Queries) AddFriendTx(ctx context.Context, userId, friendId uint32) error {
	tx, err := q.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Printf("AddFriend: begin transaction error: %v", err.Error())
		return err
	}

	sql1 := `UPDATE friends SET status = 2 WHERE user_id = $1 AND friend_id = $2`
	sql2 := `INSERT INTO friends (user_id, friend_id, status) VALUES ($1, $2, $3)`

	if _, err := tx.ExecContext(ctx, sql1, friendId, userId); err != nil {
		log.Printf("AddFriend: sql1 error: %v", err.Error())
		return tx.Rollback()
	}
	if _, err := tx.ExecContext(ctx, sql2, userId, friendId, 2); err != nil {
		log.Printf("AddFriend: sql2 error: %v", err.Error())
		return tx.Rollback()
	}

	return tx.Commit()

}

func (q *Queries) ExistsFriend(ctx context.Context, userId, friendId uint32, status int8) int8 {

	sql := `SELECT COUNT(*) FROM friends WHERE user_id = $1 AND friend_id = $2 AND status = $3`

	var count int8
	_ = q.db.GetContext(ctx, &count, sql, userId, friendId, status)

	return count
}

func (q *Queries) GetFriend(ctx context.Context, userId, friendId uint32) (*model.Friend, error) {

	sql := `SELECT * FROM friends WHERE user_id = $1 AND friend_id = $2 AND status != 3`

	var friend model.Friend
	err := q.db.GetContext(ctx, &friend, sql, userId, friendId)

	return &friend, err

}

func (q *Queries) GetFriendAll(ctx context.Context, userId uint32) ([]model.Friend, error) {

	sql := `SELECT * FROM friends WHERE user_id = $1 AND status = 1`

	friends := []model.Friend{}
	err := q.db.SelectContext(ctx, &friends, sql, userId)

	return friends, err
}

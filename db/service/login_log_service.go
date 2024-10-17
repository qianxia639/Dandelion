package db

import (
	"Dandelion/db/model"
	"context"
	"log"
)

func (q *Queries) GetLastLoginLog(ctx context.Context, userId uint32) *model.LoginLog {

	sql := `SELECT * FROM login_logs WHERE user_id = $1 AND is_login_exp = false ORDER BY log_id DESC LIMIT 1`

	var userLoginLog model.LoginLog
	err := q.db.GetContext(ctx, &userLoginLog, sql, userId)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return &userLoginLog
}

func (q *Queries) AddLoginLog(ctx context.Context, loginLog *model.LoginLog) error {
	sql := `
	INSERT INTO login_logs (
		user_id, login_time, ip, ip_addr, is_login_exp
	) VALUES (
		$1, $2, $3, $4, $5
	)`
	_, err := q.db.QueryContext(ctx, sql,
		loginLog.UserId,
		loginLog.LoginTime,
		loginLog.Ip,
		loginLog.IpAddr,
		loginLog.IsLoginExp,
	)
	return err
}

func (q *Queries) UpdateLastLoginLog(ctx context.Context, userId uint32) error {
	sql := `
	UPDATE login_logs
	SET is_login_exp = false
	WHERE log_id = (
		SELECT log_id
		FROM login_logs
		WHERE user_id = $1 AND is_login_exp = true
		ORDER By log_id DESC
		LIMIT 1
	)`

	_, err := q.db.ExecContext(ctx, sql, userId)
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

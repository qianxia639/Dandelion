-- 用户登录记录表
CREATE TABLE IF NOT EXISTS "login_logs" (
    "log_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "login_time" TIMESTAMPTZ NOT NULL,
    "ip"  VARCHAR(45) NOT NULL,
    "ip_addr" VARCHAR(20) NOT NULL,
    "is_login_exp" BOOLEAN NOT NULL DEFAULT FALSE,
    "other" VARCHAR(16) NOT NULL DEFAULT '',
    CONSTRAINT login_log_user_user_id_fk FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

COMMENT ON COLUMN "login_logs"."log_id" IS '日志ID';

COMMENT ON COLUMN "login_logs"."user_id" IS '用户ID';

COMMENT ON COLUMN "login_logs"."login_time" IS '登录时间';

COMMENT ON COLUMN "login_logs"."ip" IS 'IP地址';

COMMENT ON COLUMN "login_logs"."ip_addr" IS 'IP所属地';

COMMENT ON COLUMN "login_logs"."is_login_exp" IS '是否登录异常,F: 无异常, T: 异常';
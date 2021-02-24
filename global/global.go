package global

import (
	sf "github.com/bwmarrin/snowflake"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

var (
	MDB   *sqlx.DB
	RDB   *redis.Client
	NODE  *sf.Node
	Trans ut.Translator
)

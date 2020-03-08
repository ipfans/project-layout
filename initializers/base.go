package initializers

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

func NewConfigure() (v *viper.Viper, err error) {
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	v = viper.GetViper()
	return
}

func NewLogger(v *viper.Viper) (logger zerolog.Logger) {
	logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	if v.GetBool("debug") {
		logger = logger.Level(zerolog.DebugLevel)
	}
	log.Logger = logger
	return
}

func NewXORM(v *viper.Viper) (db xorm.EngineInterface, err error) {
	var e *xorm.Engine
	e, err = xorm.NewEngine(v.GetString("db.type"), v.GetString("db.uri"))
	if err != nil {
		return
	}
	if !v.GetBool("debug") {
		e.SetLogger(xorm.DiscardLogger{})
	} else {
		e.SetLogger(xorm.NewSimpleLogger(os.Stdout))
	}
	db = e
	return
}

func NewRedis(v *viper.Viper) (r redis.Cmdable, err error) {
	client := redis.NewClient(&redis.Options{
		Addr: v.GetString("redis.addr"),
	})
	r = client
	err = r.Ping().Err()
	return
}

func NewHTTPService() (e *gin.Engine) {
	e = gin.Default()
	return
}

func NewGRPCService() (srv *grpc.Server) {
	srv = grpc.NewServer()
	return
}

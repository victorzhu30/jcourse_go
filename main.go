package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"jcourse_go/dal"
	"jcourse_go/rpc"
)

func Init() {
	_ = godotenv.Load()
	dal.InitRedisClient()
	dal.InitDBClient()
	rpc.InitOpenAIClient()
}

func main() {
	Init()
	r := gin.Default()
	registerRouter(r)
	_ = r.Run()
}

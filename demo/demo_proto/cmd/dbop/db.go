package main

import (
	"github.com/99n/gomall/demo/demo_proto/biz/dal"
	"github.com/99n/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/99n/gomall/demo/demo_proto/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}
	dal.Init()

	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "123456"})
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "654321")
}

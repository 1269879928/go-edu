package inits

import (
	"bufio"
	"encoding/json"
	"fmt"
	"myedu/conf"
	"os"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Gorm *gorm.DB
var Config *conf.Config
const ConfigPath = "conf/config.json"

func Init()  {
	InitMysql()
	InitConfig()
}

func InitMysql()  {
	mysqlDsn := os.Getenv("MysqlDSN")
	db, err := gorm.Open("mysql", mysqlDsn)
	if err != nil {
		fmt.Errorf("connect mysql failed, err:%v\n", err)
		return
	}
	//defer db.Close()
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(5)
	// 超时时间
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.LogMode(true)
	Gorm = db
}

func InitConfig()  {
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer configFile.Close()
	reader := bufio.NewReader(configFile)
	b,err :=reader.ReadBytes('\n')
	if err !=nil {
		panic(err.Error())
	}
	var config conf.Config
	err = json.Unmarshal(b, &config)
	fmt.Printf("config:%#v\n", config)
	if err !=nil {
		panic(err.Error())
	}
}

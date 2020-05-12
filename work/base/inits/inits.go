package inits

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-edu/conf"
	"go-edu/work/entity"
	"io/ioutil"
	"os"
	"time"
)

var Gorm *gorm.DB
var Config *conf.Config
const ConfigPath = "conf/config.json"

func Init()  {
	initMysql()
	initConfig()
	easyjsonTest()
}

func initMysql()  {
	mysqlDsn := os.Getenv("MysqlDSN")
	//fmt.Println("mysql dsn:", mysqlDsn)
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
	Migration()
	//migration.Migration()
}

func initConfig()  {
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer configFile.Close()
	//reader := bufio.NewReader(configFile)
	configByte, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err.Error())
		return
	}
	config := &conf.Config{}
	err = config.UnmarshalJSON(configByte)
	//err = json.Unmarshal(configByte, &Config)
	if err !=nil {
		panic(err.Error())
	}
	Config = config
}

func easyjsonTest()  {
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer configFile.Close()
	//reader := bufio.NewReader(configFile)
	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err.Error())
		return
	}
	s := &conf.Config{}
	err = s.UnmarshalJSON(b)
	if err != nil {
		fmt.Println("err2:", err.Error())
		return
	}

}
func Migration()  {
	Gorm.
		Set("gorm:table_options", "ENGINE=InnoDB").
		Set("gorm:table_options",  "charset=utf8mb4").
		AutoMigrate(&entity.Administrators{}, &entity.AdministratorRoles{}, &entity.AdministratorPermissions{}, &entity.Courses{},
			&entity.CourseCategories{},&entity.CourseChapter{},&entity.Videos{})
		//AutoMigrate(&entity.AdministratorsInfo{}, &entity.AdministratorRoles{},&entity.AdministratorRoleRelation{}, &entity.AdministratorPermissions{}, &entity.AdministratorRolePermissionRelation{})

}

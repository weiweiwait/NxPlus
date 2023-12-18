// package dao
//
// import (
//
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//	"gopkg.in/yaml.v2"
//	"gorm.io/driver/mysql"
//	"io/ioutil"
//	"log"
//
// )
//
//	type Config struct {
//		Database struct {
//			Host     string `yaml:"host"`
//			Port     int    `yaml:"port"`
//			Username string `yaml:"username"`
//			Password string `yaml:"password"`
//			DBName   string `yaml:"dbname"`
//		} `yaml:"database"`
//	}
//
// var (
//
//	DB *gorm.DB
//
// )
//
// // 连接数据库
//
//	func InitMysql(err error) {
//		// 读取配置文件
//		configData, err := ioutil.ReadFile("config/config.yaml")
//		if err != nil {
//			log.Fatal("无法读取配置文件:", err)
//		}
//
//		// 解析配置文件
//		var config Config
//		err = yaml.Unmarshal(configData, &config)
//		if err != nil {
//			log.Fatal("无法解析配置文件:", err)
//		}
//
//		// 构建连接字符串
//		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
//			config.Database.Username,
//			config.Database.Password,
//			config.Database.Host,
//			config.Database.Port,
//			config.Database.DBName)
//		dialector := mysql.Open(connStr)
//		// 连接到MySQL数据库
//		DB, err = gorm.Open(mysql.Open(connStr))
//		if err != nil {
//			log.Fatal("无法连接到数据库:", err)
//		}
//		DB.Close()
//
//		// 测试连接
//		err = DB.DB().Ping()
//		if err != nil {
//			log.Fatal("数据库连接错误:", err)
//		}
//
//		fmt.Println("成功连接到数据库")
//	}
//
// package dao
//
// import (
//
//	"fmt"
//	"github.com/spf13/viper"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//
// )
//
// var DB = newDB()
//
//	type MySQLConfig struct {
//		Host      string
//		Port      int
//		Username  string
//		Password  string
//		Database  string
//		Charset   string
//		ParseTime bool
//	}
//
//	func newDB() *gorm.DB {
//		//设置 Viper 的配置文件名和路径
//		viper.SetConfigName("config")
//		viper.SetConfigType("yaml")
//		viper.AddConfigPath("./config")
//
//		if err := viper.ReadInConfig(); err != nil {
//			fmt.Println("Error reading config file:", err)
//
//		}
//
//		// 使用 Viper 获取 MySQL 配置
//		var mysqlConfig MySQLConfig
//		if err := viper.UnmarshalKey("mysql", &mysqlConfig); err != nil {
//			fmt.Println("Error unmarshalling MySQL config:", err)
//		}
//
//		// 打印配置信息
//		fmt.Printf("MySQL Config: %+v\n", mysqlConfig)
//		// 初始化 Gorm 连接
//		db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v",
//			mysqlConfig.Username,
//			mysqlConfig.Password,
//			mysqlConfig.Host,
//			mysqlConfig.Port,
//			mysqlConfig.Database,
//			mysqlConfig.Charset,
//			mysqlConfig.ParseTime)))
//		if err != nil {
//			fmt.Println("Error connecting to database:", err)
//		}
//		return db
//	}
package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:Fjw20030504@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}

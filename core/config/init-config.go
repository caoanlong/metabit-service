package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"metabit-service/core/models"
	"os"
)

var GDb *gorm.DB
var GConfig Conf
var GLog *zap.Logger

func Init() {
	env := "dev"
	envs := []string{"dev", "test", "prod"}
	if len(os.Args) == 2 {
		env = os.Args[1]
		if !utils.Contains(envs, env) {
			panic(errors.New("env params error"))
		}
	}
	err := initConfig(env)
	if err != nil {
		panic(err)
	}
	GDb, err = initDB()
	if err != nil {
		panic(err)
	}

	if GDb != nil {
		//registerTables(GDb) // 初始化表
		// 程序结束前关闭数据库链接
		//db, _ := GDb.DB()
		//defer func(db *sql.DB) {
		//	err := db.Close()
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//}(db)
	}
}

func initConfig(env string) (err error) {
	configType := "yaml"
	v := viper.New()
	v.SetConfigType(configType)
	v.SetConfigFile("./conf.yaml")
	v.AddConfigPath(".")
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	configs := v.AllSettings()
	// 将conf.yml中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	// 根据配置的env读取相应的配置信息
	viper.SetConfigType(configType)
	viper.SetConfigFile("./conf." + env + ".yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	if err := viper.Unmarshal(&GConfig); err != nil {
		fmt.Println(err)
	}
	return
}

func initDB() (*gorm.DB, error) {
	m := GConfig.Mysql
	if m.Dbname == "" {
		return nil, errors.New("DB name can not be null")
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config())
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db, nil
}

func GetDB() *gorm.DB {
	return GDb
}

func registerTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Token{},
		models.Network{},
		models.SysUser{},
	)
	if err != nil {
		GLog.Error("register table failed", zap.Error(err))
	}
	GLog.Info("register table success")
}

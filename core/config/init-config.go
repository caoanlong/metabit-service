package config

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"metabit-service/core/models"
	"os"
)

func init() {
	
	err := initConfig()
	if err != nil {
		panic(err)
	}
	G_DB = initDB()
	//if G_DB != nil {
	//	registerTables(G_DB) // 初始化表
	//	// 程序结束前关闭数据库链接
	//	db, _ := G_DB.DB()
	//	defer db.Close()
	//}
}

func initConfig() (err error) {
	box := packr.New("myBox", "../")
	configType := "yaml"
	defaultConfig := box.Bytes("conf.yaml")
	v := viper.New()
	v.SetConfigType(configType)
	err = v.ReadConfig(bytes.NewReader(defaultConfig))
	if err != nil {
		return
	}
	configs := v.AllSettings()
	// 将conf.yml中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	env := os.Getenv("GO_ENV")

	// 根据配置的env读取相应的配置信息
	envConfig := box.Bytes("conf.dev.yaml")
	if env != "" {
		envConfig = box.Bytes("conf." + env + ".yaml")
	}
	viper.SetConfigType(configType)
	err = viper.ReadConfig(bytes.NewReader(envConfig))
	if err != nil {
		return
	}
	if err := viper.Unmarshal(&G_CONFIG); err != nil {
		fmt.Println(err)
	}
	return
}

func initDB() *gorm.DB {
	m := G_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func registerTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Token{},
	)
	if err != nil {
		G_LOG.Error("register table failed", zap.Error(err))
	}
	G_LOG.Info("register table success")
}

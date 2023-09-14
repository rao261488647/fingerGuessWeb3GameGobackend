package config

import (
	"fingerGuessWeb3GameGoBackend/config/setting"
	"fingerGuessWeb3GameGoBackend/internal/model"
	"fingerGuessWeb3GameGoBackend/pkg/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDbEngine(db *setting.DataBase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", db.Username, db.Pwd, db.Host, db.DbName, db.Charset, db.ParseTime)
	ndb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	return ndb, nil
}

func Migrate() error {
	err := global.DbEngine.AutoMigrate(&model.Games{})
	if err != nil {
		return err
	}
	return nil
}

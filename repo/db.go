package repo

import (
	"github.com/distanceNing/testapp/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	dsn string
	db  *gorm.DB
}

var Storage *DbInstance

func GetDefaultTestDb() {
	dbconf := conf.DbConf{"127.0.0.1:3306", "root", "DLJn@123456!"}
	_ = InitStorage(&dbconf)
}

func InitStorage(dbConf *conf.DbConf) error {
	dbi, err := NewDbInstance(dbConf)
	if err != nil {
		return err
	}
	Storage = new(DbInstance)
	Storage.db = dbi.db
	Storage.dsn = dbi.dsn
	err = Storage.db.AutoMigrate((*ImageInfo)(nil), &UserInfo{}, &ArticleInfo{})
	return err
}

func NewDbInstance(dbConf *conf.DbConf) (*DbInstance, error) {
	dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Addr +
		")/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Println("failed to connect database" + err.Error())
		return nil, err
	}
	return &DbInstance{dsn, db}, nil
}

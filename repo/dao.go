package repo

import (
	"github.com/distanceNing/testapp/common"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserInfo struct {
	UserId       string `gorm:"primaryKey"`
	NickName     string
	UserType     int
	UserPassword string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ArticleInfo struct {
	Id        int `gorm:"primaryKey:autoIncrement"`
	Title     string
	Content   string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func QueryUserInfo(userId string) (common.Status, UserInfo) {
	var userInfo UserInfo
	status := common.NewStatus()
	res := Storage.db.Where(&UserInfo{UserId: userId}).First(&userInfo)
	if res.Error == gorm.ErrRecordNotFound {
		status.Set(common.ErrUserNotExist, "user not exist")
	} else if res.Error != nil {
		status.Set(common.ErrSystem, "query user failed")
	}
	return status, userInfo
}

func CreateObject(obj interface{}) common.Status {
	status := common.NewStatus()
	res := Storage.db.Create(obj)
	if res.RowsAffected == 0 {
		status.Set(common.ErrDbDupKey, "insert dup key")
		return status
	}
	return status
}

func UpdateObject(cond interface{}, updateField interface{}) common.Status {
	status := common.NewStatus()
	res := Storage.db.Model(cond).Updates(updateField)
	if res.RowsAffected == 0 {
		status.Set(common.ErrNoAffected, "db update op affected 0 row")
		return status
	}
	return status

}

type UserSession struct {
	UserId    string `gorm:"primarykey"`
	TokenId   string
	CreatedAt time.Time
}

func DeleteSession(userId string) common.Status {
	status := common.NewStatus()
	res := Storage.db.Delete(&UserSession{}, userId)
	if res.Error != nil {
		log.Println(res.Error.Error())
		status.Set(common.ErrSystem, "db delete failed")
	}
	return status
}

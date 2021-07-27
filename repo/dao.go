package repo

import (
	"github.com/distanceNing/testapp/common"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserInfo struct {
	UserId       string `gorm:"primarykey"`
	UserType     int
	UserPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func CreateUser(userId string, userType int, password string) common.Status {
	status := common.NewStatus()
	res := Storage.db.Create(&UserInfo{userId, userType, password, time.Now(), time.Now()})
	if res.RowsAffected == 0 {
		status.Set(common.ErrDbDupKey, "insert dup key")
		return status
	}
	return status
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

type UserSession struct {
	UserId    string `gorm:"primarykey"`
	TokenId   string
	CreatedAt time.Time
}

func CreateSession(userId string, token string) common.Status {
	status := common.NewStatus()
	res := Storage.db.Create(&UserSession{userId, token, time.Now()})
	if res.RowsAffected == 0 {
		status.Set(common.ErrDbDupKey, "insert dup key")
		return status
	}
	return status
}

func QuerySessionToken(userId string) (common.Status, UserSession) {
	var session UserSession
	status := common.NewStatus()
	res := Storage.db.Where(&UserSession{UserId: userId}).First(&session)
	if res.Error != nil {
		status.Set(common.ErrSystem, "query session failed")
		return status, session
	}
	return status, session
}

func UpdateSessionToken(userId string, token string) common.Status {
	status := common.NewStatus()
	res := Storage.db.Save(&UserSession{userId, token, time.Now()})
	if res.RowsAffected == 0 {
		status.Set(common.ErrDbDupKey, "insert dup key")
		return status
	}
	return status
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

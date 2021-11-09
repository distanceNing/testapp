package repo

import (
	"github.com/distanceNing/testapp/src/common/errcode"
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
	Id        int       `gorm:"primaryKey:autoIncrement"`
	ChannelId int       `json:"channel_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Images    string    `json:"images"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ImageInfo struct {
	Id       int64 `gorm:"primaryKey"`
	BelongTo string
	Url      string
}

// 分页查询
func QueryObjectByPage(cond interface{}, objs interface{}, pageCount int, pageNum int) error {
	res := Storage.db.Limit(pageCount).Offset(pageCount * (pageNum - 1)).Where(cond).Find(objs)
	if res.Error == gorm.ErrRecordNotFound {
		return errcode.New(errcode.ErrRecordNotExist, "record not exist")
	} else if res.Error != nil {
		return errcode.New(errcode.ErrSystem, "query record failed")
	}
	return nil
}

func QueryObjectCount(cond interface{}, count *int64) error {
	res := Storage.db.Model(cond).Where(cond).Count(count)
	if res.Error == gorm.ErrRecordNotFound {
		return errcode.New(errcode.ErrRecordNotExist, "record not exist")
	} else if res.Error != nil {
		log.Println(res.Error.Error())
		return errcode.New(errcode.ErrSystem, "query record failed")
	}
	return nil
}

func QueryUserInfo(userId string) (error, UserInfo) {
	var userInfo UserInfo
	res := Storage.db.Where(&UserInfo{UserId: userId}).First(&userInfo)
	if res.Error == gorm.ErrRecordNotFound {
		return errcode.New(errcode.ErrUserNotExist, "user not exist"), userInfo
	} else if res.Error != nil {
		return errcode.New(errcode.ErrSystem, "query user failed"), userInfo
	}
	return nil, userInfo
}

func QueryObject(cond interface{}, obj interface{}) error {
	res := Storage.db.Where(cond).First(obj)
	if res.Error == gorm.ErrRecordNotFound {
		return errcode.New(errcode.ErrRecordNotExist, "record not exist")
	} else if res.Error != nil {
		return errcode.New(errcode.ErrSystem, "query record failed")
	}
	return nil
}

func CreateObject(obj interface{}) error {
	res := Storage.db.Create(obj)
	if res.RowsAffected == 0 {
		return errcode.New(errcode.ErrDbDupKey, "insert dup key")
	}
	return nil
}

func UpdateObject(cond interface{}, updateField interface{}) error {
	res := Storage.db.Model(cond).Updates(updateField)
	if res.RowsAffected == 0 {
		return errcode.New(errcode.ErrNoAffected, "db update op affected 0 row")
	}
	return nil
}

func DeleteObject(cond interface{}) error {

	res := Storage.db.Model(cond).Delete(cond)
	if res.RowsAffected == 0 {
		return errcode.New(errcode.ErrNoAffected, "db delete op affected 0 row")
	}
	return nil
}

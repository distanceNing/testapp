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
	Id        int       `gorm:"primaryKey:autoIncrement"`
	ChannelId int       `json:"channel_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Images    string    `json:"images"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentInfo struct {
	Id          int    `gorm:"primaryKey:autoIncrement"`
	BelongTo    int    `gorm:"index"` // 属于那个作品
	PublisherId string `gorm:"index"`
	Content     string
	Status      int
	CreatedAt   time.Time
}

type CommentReply struct {
	ParentId    int    `gorm:"primaryKey"` // 父评论id
	PublisherId string `gorm:"index"`
	Content     string
	Status      int
	CreatedAt   time.Time
}

type ImageInfo struct {
	Id       int64 `gorm:"primaryKey"`
	BelongTo string
	Url      string
}

// 分页查询
func QueryObjectByPage(cond interface{}, objs interface{}, pageCount int, pageNum int) common.Status {
	status := common.NewStatus()
	res := Storage.db.Limit(pageCount).Offset(pageCount * (pageNum - 1)).Where(cond).Find(objs)
	if res.Error == gorm.ErrRecordNotFound {
		status.Set(common.ErrRecordNotExist, "record not exist")
	} else if res.Error != nil {
		status.Set(common.ErrSystem, "query record failed")
	}
	return status
}

func QueryObjectCount(cond interface{}, count *int64) common.Status {
	status := common.NewStatus()
	res := Storage.db.Model(cond).Where(cond).Count(count)
	if res.Error == gorm.ErrRecordNotFound {
		status.Set(common.ErrRecordNotExist, "record not exist")
	} else if res.Error != nil {
		log.Println(res.Error.Error())
		status.Set(common.ErrSystem, "query record failed")
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

func QueryObject(cond interface{}, obj interface{}) common.Status {
	status := common.NewStatus()
	res := Storage.db.Where(cond).First(obj)
	if res.Error == gorm.ErrRecordNotFound {
		status.Set(common.ErrRecordNotExist, "record not exist")
	} else if res.Error != nil {
		status.Set(common.ErrSystem, "query record failed")
	}
	return status
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

func DeleteObject(cond interface{}) common.Status {
	status := common.NewStatus()
	res := Storage.db.Model(cond).Delete(cond)
	if res.RowsAffected == 0 {
		status.Set(common.ErrNoAffected, "db delete op affected 0 row")
		return status
	}
	return status
}

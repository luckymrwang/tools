package models

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DH_Corp struct {
	Id         string    `json:"id" gorm:"column:id"`
	ObjectId   string    `json:"objectId" gorm:"column:object_id"`
	Name       string    `json:"name" gorm:"column:name"`
	Email      string    `json:"email" gorm:"column:email"`
	Mobile     string    `json:"mobile" gorm:"column:mobile"`
	Vcode      string    `json:"vcode" gorm:"column:vcode"`
	ConnectId  string    `json:"connectId" gorm:"column:connect_id"`
	Status     string    `json:"status" gorm:"column:status"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`
}

type User struct {
	ID           uint           `json:"id" gorm:"column:id"`
	Name         string         `json:"name" gorm:"column:name"`
	Email        *string        `json:"email" gorm:"column:name"`
	Age          uint8          `json:"age" gorm:"column:age"`
	Birthday     *time.Time     `json:"birthday" gorm:"column:birthday"`
	MemberNumber sql.NullString `json:"memberNumber" gorm:"column:memberNumber"`
	ActivatedAt  sql.NullTime   `json:"activatedAt" gorm:"column:activatedAt"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
}

func (m *DH_Corp) Table(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 {
		return tx[0].Table("dh_corp")
	}
	return DB.Table("dh_corp")
}

func (m *DH_Corp) TableName() string {
	return "dh_corp"
}

func (m *DH_Corp) OrderPager(filter map[string]interface{}, page, pageSize int, orderBy, order string) (count int64, list []*DH_Corp, err error) {
	db := m.Table().Where(filter).Count(&count).Order(fmt.Sprintf("%v %v", orderBy, order))
	if pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err = db.Find(&list).Error
	return
}

func (m *DH_Corp) Delete(dhc *DH_Corp, tx ...*gorm.DB) error {
	if err := m.Table(tx...).Delete(dhc).Error; err != nil {
		return err
	}
	return nil
}

func (m *DH_Corp) Update(dhc *DH_Corp, value interface{}, tx ...*gorm.DB) error {
	if err := m.Table(tx...).Model(dhc).Update("name", value).Error; err != nil {
		return err
	}
	return fmt.Errorf("ccccc")
}

func (m *DH_Corp) TM(dhc *DH_Corp) error {
	// 安全的，会被转义
	if err := m.Table().Debug().Where("name = ?", "cl;drop table users;").First(dhc).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

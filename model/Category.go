package model

import (
	"boKe/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"varchar(20);not null" json:"name"`
}

//查询分类是否存在

func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCESS //200
}

//新增分类

func CreateCategory(data *Category) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Debug().Create(&data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

//查询分类列表

func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

//	编辑分类信息

func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Debug().Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//	删除分类

func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

package model

import (
	"boKe/utils/errmsg"
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

//新增文章

func CreateArt(data *Article) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Debug().Create(&data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

//  查询分类下所有文章

func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_CAT_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

//  查询单个文章

func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

//查询文字列表

func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Preload("Category").Limit(pageSize).Offset(offset).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

//	编辑文章

func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Debug().Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//	删除文章

func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

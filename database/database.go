package database

import (
	"fmt"
	"my-personal-website/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseService interface {
	All() models.PersonalInfo
	Update(info models.PersonalInfo) error
	FindBlog() []models.Article
	ShowArticle(id int) models.Article
	AddArticle(a models.Article) error
	DeleteArticle(id int) error
}

type database struct {
	connection *gorm.DB
}

func NewDatabaseService() DatabaseService {
	dsn := "root:password@tcp(mysqlc:3306)/website?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.PersonalInfo{})
	return &database{
		connection: db,
	}
}

func (db *database) All() models.PersonalInfo {
	var pi models.PersonalInfo
	db.connection.First(&pi, 1)
	return pi
}

func (db *database) Update(info models.PersonalInfo) error {
	return db.connection.Model(&models.PersonalInfo{}).Where("id = ?", info.ID).Updates(info).Error
}

func (db *database) FindBlog() []models.Article {
	var blogs []models.Article
	db.connection.Find(&blogs)
	return blogs
}

func (db *database) ShowArticle(id int) models.Article {
	var article models.Article
	db.connection.First(&article, id)
	return article
}

func (db *database) AddArticle(a models.Article) error {
	err := db.connection.Create(&a).Error
	return err
}

func (db *database) DeleteArticle(id int) error {
	fmt.Println("*****id:", id)

	err := db.connection.Delete(&models.Article{}, id).Error
	return err
}

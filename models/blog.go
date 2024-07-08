package models

type Article struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}

type PersonalInfo struct {
	ID          uint   `json:"id" gorm:"primaryKey" form:"id"`
	Title       string `json:"title" gorm:"size:255" form:"title"`
	Description string `json:"description" gorm:"type:text" form:"description"`
	Interest    string `json:"interest" gorm:"type:text" form:"interest"`
	Contact     string `json:"contact" gorm:"size:255" form:"contact"`
}

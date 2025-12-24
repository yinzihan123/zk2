package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Mobile   string `gorm:"type:varchar(30)" json:"Mobile"`
	Password string `gorm:"type:char(32)" json:"Password"`
	Salt     string `gorm:"type:varchar(30)" json:"Salt"`
}
type Banner struct {
	gorm.Model
	Pic   string `gorm:"type:varchar(300)" json:"Pic"`
	Title string `gorm:"type:varchar(50)" json:"Title"`
}
type Good struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(300)" json:"Name"`
	Pic           string  `gorm:"type:varchar(300)" json:"Pic"`
	Price         float64 `gorm:"type:decimal(10,2)" json:"Price"`
	OriginalPrice float64 `gorm:"type:decimal(10,2)" json:"OriginalPrice"`
	SalesVolume   float64 `gorm:"type:decimal(10,2)" json:"SalesVolume"`
	ShopName      string  `gorm:"type:varchar(50)" json:"ShopName"`
	Inventory     int     `gorm:"type:int(11)" json:"Inventory"`
}

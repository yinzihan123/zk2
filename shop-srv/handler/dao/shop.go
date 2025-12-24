package dao

import (
	"gorm.io/gorm"
	"shop-srv/handler/models"
)

type ShopDao struct {
}

func NewShopDao() *ShopDao {
	return &ShopDao{}
}
func (s *ShopDao) FindUserByMobile(db *gorm.DB, mobile string) (user *models.User, err error) {
	err = db.Where("mobile=?", mobile).Find(&user).Error
	return
}

func (s *ShopDao) UserAdd(db *gorm.DB, user *models.User) (*models.User, error) {
	err := db.Create(&user).Error
	return user, err
}

func (s *ShopDao) FindGoodList(db *gorm.DB) (good []*models.Good, err error) {
	err = db.Find(&good).Error
	return
}
func (s *ShopDao) FindBannerList(db *gorm.DB) (banner []*models.Banner, err error) {
	err = db.Find(&banner).Error
	return
}

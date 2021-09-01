package models

import (
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	ID    int
	Name  string
	Address string
	Phone string
	Email string
}

func CreateSeller(db *gorm.DB, Seller *Seller) (err error) {
	err = db.Create(Seller).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSellers(db *gorm.DB, Seller *[]Seller) (err error) {
	err = db.Find(Seller).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSeller(db *gorm.DB, Seller *Seller, id string) (err error) {
	err = db.Where("id = ?", id).First(Seller).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateSeller(db *gorm.DB, Seller *Seller) (err error) {
	db.Save(Seller)
	return nil
}

func DeleteSeller(db *gorm.DB, Seller *Seller, id string) (err error) {
	db.Where("id = ?", id).Delete(Seller)
	return nil
}
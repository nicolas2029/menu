package controller

import (
	"menu/model"
	"menu/storage"
)

// GetProduct return a product by ID
func GetProduct(id uint) (model.Product, error) {
	p := model.Product{}
	err := storage.DB().First(&p, id).Error
	return p, err
}

// GetProducts return all products
func GetAllProduct() ([]model.Product, error) {
	ps := make([]model.Product, 0)
	r := storage.DB().Find(&ps)
	return ps, r.Error
}

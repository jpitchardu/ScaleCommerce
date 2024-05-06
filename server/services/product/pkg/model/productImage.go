package model

import "database/sql"

type ProductImage struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"productId"`
	URL       string `json:"url"`
}

type ProductImageModel interface {
	CreateProductImage(productImage *ProductImage) (int64, error)
	GetProductImage(id int64) (*ProductImage, error)
	UpdateProductImage(productImage *ProductImage) (int64, error)
	DeleteProductImage(id int64) error
}

type productImageModel struct{ DB *sql.DB }

// CreateProductImage implements ProductImageModel.
func (p *productImageModel) CreateProductImage(productImage *ProductImage) (int64, error) {
	panic("unimplemented")
}

// DeleteProductImage implements ProductImageModel.
func (p *productImageModel) DeleteProductImage(id int64) error {
	panic("unimplemented")
}

// GetProductImage implements ProductImageModel.
func (p *productImageModel) GetProductImage(id int64) (*ProductImage, error) {
	panic("unimplemented")
}

// UpdateProductImage implements ProductImageModel.
func (p *productImageModel) UpdateProductImage(productImage *ProductImage) (int64, error) {
	panic("unimplemented")
}

func NewProductImageModel(db *sql.DB) ProductImageModel { return &productImageModel{DB: db} }

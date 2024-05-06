package model

import "database/sql"

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"productId"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
}

type ProductModel interface {
	CreateProduct(product *Product) (int64, error)
	GetProduct(id int64) (*Product, error)
	UpdateProduct(product *Product) (int64, error)
	DeleteProduct(id int64) error
}

type productModel struct{ DB *sql.DB }

// CreateProduct implements ProductModel.
func (p *productModel) CreateProduct(product *Product) (int64, error) {
	panic("unimplemented")
}

// DeleteProduct implements ProductModel.
func (p *productModel) DeleteProduct(id int64) error {
	panic("unimplemented")
}

// GetProduct implements ProductModel.
func (p *productModel) GetProduct(id int64) (*Product, error) {
	panic("unimplemented")
}

// UpdateProduct implements ProductModel.
func (p *productModel) UpdateProduct(product *Product) (int64, error) {
	panic("unimplemented")
}

func NewProductModel(db *sql.DB) ProductModel { return &productModel{DB: db} }

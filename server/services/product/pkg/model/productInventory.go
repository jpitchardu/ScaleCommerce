package model

import "database/sql"

type ProductInventory struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

type ProductInventoryModel interface {
	CreateProductInventory(productInventory *ProductInventory) (int64, error)
	GetProductInventory(id int64) (*ProductInventory, error)
	UpdateProductInventory(productInventory *ProductInventory) (int64, error)
	DeleteProductInventory(id int64) error
}

type productInventoryModel struct{ DB *sql.DB }

// CreateProductInventory implements ProductInventoryModel.
func (p *productInventoryModel) CreateProductInventory(productInventory *ProductInventory) (int64, error) {
	panic("unimplemented")
}

// DeleteProductInventory implements ProductInventoryModel.
func (p *productInventoryModel) DeleteProductInventory(id int64) error {
	panic("unimplemented")
}

// GetProductInventory implements ProductInventoryModel.
func (p *productInventoryModel) GetProductInventory(id int64) (*ProductInventory, error) {
	panic("unimplemented")
}

// UpdateProductInventory implements ProductInventoryModel.
func (p *productInventoryModel) UpdateProductInventory(productInventory *ProductInventory) (int64, error) {
	panic("unimplemented")
}

func NewProductInventoryModel(db *sql.DB) ProductInventoryModel {
	return &productInventoryModel{DB: db}
}

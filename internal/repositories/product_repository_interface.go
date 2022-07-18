package repositories

import "mockery/internal/models"

type ProductRepositoryInterface interface {
	Add(product models.Product) error
}

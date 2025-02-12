package productrepository

import (
	"github.com/lukasrvs/clean-architecture-golang/adapter/postgres"
	"github.com/lukasrvs/clean-architecture-golang/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}

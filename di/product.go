package di

import (
	"github.com/lukasrvs/clean-architecture-golang/adapter/http/productservice"
	"github.com/lukasrvs/clean-architecture-golang/adapter/postgres"
	"github.com/lukasrvs/clean-architecture-golang/adapter/postgres/productrepository"
	"github.com/lukasrvs/clean-architecture-golang/core/domain"
	"github.com/lukasrvs/clean-architecture-golang/core/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}

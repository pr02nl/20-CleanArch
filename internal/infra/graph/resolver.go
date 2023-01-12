package graph

import "github.com/pr02nl/20-CleanArch/internal/usecases"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecases.CreateOrderUseCase
	ListOrdersUseCase  usecases.ListOrdersUseCase
}

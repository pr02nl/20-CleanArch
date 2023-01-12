package usecases

import "github.com/pr02nl/20-CleanArch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (g *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := g.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var outputOrders []OrderOutputDTO
	for _, order := range orders {
		outputOrders = append(outputOrders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return outputOrders, nil
}

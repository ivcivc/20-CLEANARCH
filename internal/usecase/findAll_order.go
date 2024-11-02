package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type Order struct {
	ID         string // ou outro tipo, dependendo da sua implementação
	Price      float64
	Tax        float64
	FinalPrice float64
}

type FindAllOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindAllOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *FindAllOrderUseCase {
	return &FindAllOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (f *FindAllOrderUseCase) Execute() ([]entity.Order, error) {
	orders, err := f.OrderRepository.FindAll()

	if err != nil {
		return nil, err
	}
	return orders, nil
}

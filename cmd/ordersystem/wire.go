//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/pr02nl/20-CleanArch/internal/entity"
	"github.com/pr02nl/20-CleanArch/internal/event"
	"github.com/pr02nl/20-CleanArch/internal/infra/database"
	"github.com/pr02nl/20-CleanArch/internal/infra/web"
	"github.com/pr02nl/20-CleanArch/internal/usecases"
	"github.com/pr02nl/20-CleanArch/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecases.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecases.NewCreateOrderUseCase,
	)
	return &usecases.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB) *usecases.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecases.NewListOrdersUseCase,
	)
	return &usecases.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

package web

import (
	"encoding/json"
	"net/http"

	"github.com/pr02nl/20-CleanArch/internal/entity"
	"github.com/pr02nl/20-CleanArch/internal/usecases"
	"github.com/pr02nl/20-CleanArch/pkg/events"
)

type WebOrderHandler struct {
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
	EventDispatcher   events.EventDispatcherInterface
}

func NewWebOrderHandler(
	orderRepository entity.OrderRepositoryInterface,
	orderCreatedEvent events.EventInterface,
	eventDispatcher events.EventDispatcherInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository:   orderRepository,
		OrderCreatedEvent: orderCreatedEvent,
		EventDispatcher:   eventDispatcher,
	}
}

func (h *WebOrderHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listOrders(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (h *WebOrderHandler) create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createOrder := usecases.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) listOrders(w http.ResponseWriter, r *http.Request) {
	listOrders := usecases.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	order "github.com/postech-soat2-grupo16/pedidos-api/adapters/order"
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"github.com/postech-soat2-grupo16/pedidos-api/interfaces"
	"github.com/postech-soat2-grupo16/pedidos-api/util"
	"net/http"
)

type OrderController struct {
	useCase interfaces.OrderUseCase
}

func NewOrderController(useCase interfaces.OrderUseCase, r *chi.Mux) *OrderController {
	controller := OrderController{useCase: useCase}
	r.Route("/pedidos", func(r chi.Router) {
		r.Get("/", controller.GetAll)
		r.Post("/", controller.Create)
		r.Get("/{id}", controller.GetByID)
		r.Put("/{id}", controller.Update)
		r.Delete("/{id}", controller.Delete)
		r.Patch("/{id}", controller.PatchOrderStatus)
		r.Get("/healthcheck", controller.Ping)
	})
	return &controller
}

// @Summary	health check endpoint
//
// @Tags		Orders
//
// @ID			health-check
// @Success	200
// @Router		/pedidos/healtcheck [get]
func (c *OrderController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetAll @Summary	Gets all orders by filters
//
// @Tags		Orders
// @ID			get-all-orders
// @Produce	json
//
// @Param       client_id  query       string  false   "Optional Filter by client_id"
// @Param       status  query       string  false   "Optional Filter by order status"
//
// @Success	200	{object}	order.Order
// @Failure	500
// @Router		/pedidos [get]
func (c *OrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	status := r.URL.Query().Get("status")

	ordersFetched, err := c.useCase.List(clientID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var orders []*order.Order

	if ordersFetched != nil {
		for _, orderFetched := range *ordersFetched {
			orders = append(orders, order.FromUseCaseEntity(&orderFetched))
		}
		json.NewEncoder(w).Encode(orders)
	}

	json.NewEncoder(w).Encode([]*order.Order{})
}

// @Summary	Gets an order by ID
//
// @Tags		Orders
//
// @ID			get-order-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	order.Order
// @Failure	404
// @Router		/pedidos/{id} [get]
func (c *OrderController) GetByID(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")
	if orderID == "" {
		http.Error(w, util.NewErrorDomain("order_id URL Param is missing").Error(), http.StatusBadRequest)
		return
	}

	orderFetched, err := c.useCase.GetByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if orderFetched == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order.FromUseCaseEntity(orderFetched))
}

// @Summary	New order
//
// @Tags		Orders
//
// @ID			create-order
// @Produce	json
// @Param		data	body		order.Order	true	"Order payload"
// @Success	200		{object}	order.Order
// @Failure	400
// @Router		/pedidos [post]
func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var orderModel order.Order
	err := json.NewDecoder(r.Body).Decode(&orderModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.NewErrorDomain("Error parsing request body"))
		return
	}
	orderCreated, err := c.useCase.Create(orderModel.ToUseCaseEntity())
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order.FromUseCaseEntity(orderCreated))
}

// @Summary	Updates an order
//
// @Tags		Orders
//
// @ID			update-order
// @Produce	json
// @Param		id		path		string	true	"Order ID"
// @Param		data	body		order.Order	true	"Order payload"
// @Success	200		{object}	order.Order
// @Failure	404
// @Failure	400
// @Router		/pedidos/{id} [put]
func (c *OrderController) Update(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")
	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.NewErrorDomain("id URL Param is missing"))
		return
	}

	var o order.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.NewErrorDomain("Error parsing request body"))
		return
	}

	order, err := c.useCase.Update(orderID, o.ToUseCaseEntity())
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// @Summary	Patches order's status
//
// @Tags		Orders
//
// @ID			update-status-order
// @Produce	json
// @Param		id		path		string	true	"Order ID"
// @Param		data	body		order.Order	true	"Order with updated status"
// @Success	200		{object}	order.Order
// @Failure	404
// @Failure	400
// @Router		/pedidos/{id} [patch]
func (c *OrderController) PatchOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")
	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.NewErrorDomain("id URL Param is missing"))
		return
	}

	var o order.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.NewErrorDomain("Error parsing request body"))
		return
	}

	order, err := c.useCase.UpdateOrderStatus(orderID, entities.Status(o.Status))
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if order == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// @Summary	Deletes an order by ID
//
// @Tags		Orders
//
// @ID			delete-order-by-id
// @Produce	json
// @Param		id	path	string	true	"Order ID"
// @Success	204
// @Failure	500
// @Router		/pedidos/{id} [delete]
func (c *OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")
	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(util.NewErrorDomain("id URL Param is missing"))
		return
	}

	err := c.useCase.Delete(orderID)
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

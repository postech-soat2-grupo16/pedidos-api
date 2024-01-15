package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	order "github.com/postech-soat2-grupo16/pedidos-api/adapters/order"
	"github.com/postech-soat2-grupo16/pedidos-api/interfaces"
	"github.com/postech-soat2-grupo16/pedidos-api/util"
)

type OrderController struct {
	useCase interfaces.OrderUseCase
}

func NewOrderController(useCase interfaces.OrderUseCase, r *chi.Mux) *OrderController {
	controller := OrderController{useCase: useCase}
	r.Route("/orders", func(r chi.Router) {
		r.Get("/", controller.GetAll)
		r.Post("/", controller.Create)
		r.Get("/{id}", controller.GetByID)
		r.Put("/{id}", controller.Update)
		r.Delete("/{id}", controller.Delete)
		r.Patch("/{id}", controller.PatchOrderStatus)
	})
	return &controller
}

// @Summary	Gets all orders by filters
//
// @Tags		Orders
//
// @ID			get-all-orders
// @Produce	json
// @Param       client_id  query       string  false   "Optional Filter by client_id"
// @Param       status  query       string  false   "Optional Filter by order status"
// @Success	200	{object}	order.Order
// @Failure	500
// @Router		/orders [get]
func (c *OrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	orders, err := c.useCase.List(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(orders)
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
// @Router		/orders/{id} [get]
func (c *OrderController) GetByID(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")
	if orderId == "" {
		http.Error(w, util.NewErrorDomain("order_id URL Param is missing").Error(), http.StatusBadRequest)
		return
	}

	orderFetched, err := c.useCase.GetByID(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if orderFetched == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order.OrderFromEntity(orderFetched))
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
// @Router		/orders [post]
func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var orderModel order.Order
	err := json.NewDecoder(r.Body).Decode(&orderModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	json.NewEncoder(w).Encode(order.OrderFromEntity(orderCreated))
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
// @Router		/orders/{id} [put]
func (c *OrderController) Update(w http.ResponseWriter, r *http.Request) {
	var p order.Order
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := c.useCase.Update(string(id), p.ToUseCaseEntity())
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
// @Router		/orders/{id} [patch]
func (c *OrderController) PatchOrderStatus(w http.ResponseWriter, r *http.Request) {
	var p order.Order
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := c.useCase.UpdateOrderStatus(string(id), p.Status)
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
// @Router		/orders/{id} [delete]
func (c *OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.useCase.Delete(string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

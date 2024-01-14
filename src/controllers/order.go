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
		r.Patch("/{id}", controller.PatchPedidoStatus)
	})
	return &controller
}

// @Summary	Get all orders
//
// @Tags		Orders
//
// @ID			get-all-orders
// @Produce	json
// @Param       status  query       string  false   "Optional Filter by Status"
// @Success	200	{object}	order.Pedido
// @Failure	500
// @Router		/pedidos [get]
func (c *OrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	orders, err := c.useCase.List(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(orders)
}

// @Summary	Get a order by ID
//
// @Tags		Orders
//
// @ID			get-order-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	order.Pedido
// @Failure	404
// @Router		/pedidos/{id} [get]
func (c *OrderController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pedido, err := c.useCase.GetByID(string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if pedido == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pedido)
}

// @Summary	New order
//
// @Tags		Orders
//
// @ID			create-order
// @Produce	json
// @Param		data	body		order.Pedido	true	"Order data"
// @Success	200		{object}	order.Pedido
// @Failure	400
// @Router		/pedidos [post]
func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var p order.Order
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := c.useCase.Create(p.ToEntity())
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
	json.NewEncoder(w).Encode(order)
}

// @Summary	Update a order
//
// @Tags		Orders
//
// @ID			update-order
// @Produce	json
// @Param		id		path		string	true	"Order ID"
// @Param		data	body		order.Pedido	true	"Order data"
// @Success	200		{object}	order.Pedido
// @Failure	404
// @Failure	400
// @Router		/pedidos/{id} [put]
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
	order, err := c.useCase.Update(string(id), p.ToEntity())
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

// @Summary	Patch status of a order
//
// @Tags		Orders
//
// @ID			update-status-order
// @Produce	json
// @Param		id		path		string	true	"Order ID"
// @Param		data	body		order.Pedido	true	"Pedido with updated status"
// @Success	200		{object}	order.Pedido
// @Failure	404
// @Failure	400
// @Router		/pedidos/{id} [patch]
func (c *OrderController) PatchPedidoStatus(w http.ResponseWriter, r *http.Request) {
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

// @Summary	Delete a order by ID
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

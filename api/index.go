package api

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/postech-soat2-grupo16/pedidos-api/controllers"
	"github.com/postech-soat2-grupo16/pedidos-api/external"
	og "github.com/postech-soat2-grupo16/pedidos-api/gateways/db/order"
	"github.com/postech-soat2-grupo16/pedidos-api/usecases/order"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupDB() *dynamodb.DynamoDB {
	return external.GetDynamoDbClient()
}

func SetupRouter(db *dynamodb.DynamoDB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db)

	return r
}

func mapRoutes(r *chi.Mux, db *dynamodb.DynamoDB) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Gateways
	orderGateway := og.NewGateway(db)
	// Use cases
	orderUseCase := order.NewUseCase(orderGateway)
	// Handlers
	_ = controllers.NewOrderController(orderUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

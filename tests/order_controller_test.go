package tests

import (
	"errors"
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/postech-soat2-grupo16/pedidos-api/controllers"
	"github.com/postech-soat2-grupo16/pedidos-api/interfaces/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	// Use readable error values for better test readability
	errUsecaseFailure  = errors.New("ErrUsecaseFailed")
	errUsecaseNotFound = errors.New("ErrUsecaseNotFound")
)

func TestGetAll_Error(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("List", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pedidos", nil)

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	// Assuming the expected status code is 500 for the kind of test case
	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestGetByID_Error(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("GetByID", mock.Anything).Return(nil, errUsecaseNotFound)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pedidos/1", nil)

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Not Found response is expected")
}

func TestCreate_ErrorParse(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("Create", mock.Anything).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	badJSON := `{"invalid json`
	req, _ := http.NewRequest("POST", "/pedidos", strings.NewReader(badJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Internal Server Error response is expected")
}

func TestCreate_ErrorUsecase(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("Create", mock.Anything).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	okJSON := `{}`
	req, _ := http.NewRequest("POST", "/pedidos", strings.NewReader(okJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestPUT_ErrorParse(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("Create", mock.Anything).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	badJSON := `{"invalid json`
	req, _ := http.NewRequest("PUT", "/pedidos/1", strings.NewReader(badJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Internal Server Error response is expected")
}

func TestPUT_ErrorUsecase(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("Update", mock.Anything, mock.Anything).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	okJSON := `{}`
	req, _ := http.NewRequest("PUT", "/pedidos/1", strings.NewReader(okJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestPATCH_ErrorParse(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("UpdateOrderStatus", mock.Anything).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	badJSON := `{"invalid json`
	req, _ := http.NewRequest("PATCH", "/pedidos/1", strings.NewReader(badJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Internal Server Error response is expected")
}

func TestPATCH_ErrorUsecase(t *testing.T) {
	useCase := new(mocks.OrderUseCase)
	useCase.On("UpdateOrderStatus", mock.Anything, mock.Anything).Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	okJSON := `{}`
	req, _ := http.NewRequest("PATCH", "/pedidos/1", strings.NewReader(okJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestGET_Success(t *testing.T) {
	var orderedItems []entities.OrderedItem
	var newOrder *entities.Order

	orderedItems = append(orderedItems, entities.OrderedItem{
		ItemID:      "1",
		Price:       10,
		Quantity:    1,
		Name:        "nome",
		Category:    "categoria",
		Description: "descricao",
	})
	newOrder = &entities.Order{
		OrderID:      "1",
		ClientID:     "123",
		Status:       "CRIADO",
		OrderedItems: orderedItems,
		Notes:        "nota",
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}

	useCase := new(mocks.OrderUseCase)
	useCase.On("GetByID", mock.Anything).Return(newOrder, nil)

	res := httptest.NewRecorder()
	okJSON := `{}`
	req, _ := http.NewRequest("GET", "/pedidos/1", strings.NewReader(okJSON))

	c := chi.NewRouter()
	controllers.NewOrderController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "")
}

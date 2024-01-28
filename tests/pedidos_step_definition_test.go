package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/postech-soat2-grupo16/pedidos-api/entities"
)

func parameterID(pedidoID string) error {
	inputs.pedidoID = pedidoID
	return nil
}

func requestPOSTPedido() error {
	pedidoItem := entities.Order{OrderID: inputs.pedidoID}
	body, err := json.Marshal(pedidoItem)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/pedidos", baseURL), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func statusCodeShouldBe(statusCode int) error {
	if inputs.statusCode != statusCode {
		return fmt.Errorf("expected status code %d, got %d", statusCode, inputs.statusCode)
	}
	return nil
}

func getFirstOrderID() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pedidos", baseURL), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	var orders []entities.Order
	err = json.NewDecoder(res.Body).Decode(&orders)
	if err != nil {
		return err
	}

	inputs.firstOrder = orders[0]
	inputs.pedidoID = orders[0].OrderID
	return nil
}

func parameterClientID(arg1 int) error {
	inputs.clientID = arg1
	return nil
}

func requestGETHealthcheck() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pedidos/healthcheck", baseURL), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func requestGETPedidoById() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pedidos/%s", baseURL, inputs.pedidoID), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	inputs.body = res.Body
	return nil
}

func requestGETPedidoWithClientID() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pedidos?client_id=%d", baseURL, inputs.clientID), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func requestPATCHPedidoWithStatus(arg1 string) error {
	order := entities.Order{
		Status: entities.Status(arg1),
	}
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/pedidos/%s", baseURL, inputs.pedidoID), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func requestPUTPedidoWithStatus(arg1 string) error {
	order := entities.Order{
		Status: entities.Status(arg1),
	}
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/pedidos/%s", baseURL, inputs.pedidoID), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func unknownOrderID() error {
	inputs.pedidoID = "9999"
	return nil
}

func parameterStatus(arg1 string) error {
	inputs.status = arg1
	return nil
}

func requestGETPedidoWithStatus() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pedidos?status=%s", baseURL, inputs.status), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func requestDELETEPedidoById() error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/pedidos/%s", baseURL, inputs.pedidoID), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	inputs.statusCode = res.StatusCode
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Parameter ID: (\d+)$`, parameterID)
	ctx.Step(`^statusCode should be (\d+)$`, statusCodeShouldBe)
	ctx.Step(`^get first order ID$`, getFirstOrderID)
	ctx.Step(`^request GET \/healthcheck$`, requestGETHealthcheck)
	ctx.Step(`^request PATCH \/pedido with status "([^"]*)"$`, requestPATCHPedidoWithStatus)
	ctx.Step(`^request PUT \/pedido with status "([^"]*)"$`, requestPUTPedidoWithStatus)
	ctx.Step(`^unknown order ID$`, unknownOrderID)

	ctx.Step(`^Parameter ClientID: (\d+)$`, parameterClientID)
	ctx.Step(`^Parameter Status: "([^"]*)"$`, parameterStatus)
	ctx.Step(`^request GET \/pedido by id$`, requestGETPedidoById)
	ctx.Step(`^request GET \/pedido with ClientID$`, requestGETPedidoWithClientID)
	ctx.Step(`^request POST \/pedido$`, requestPOSTPedido)
	ctx.Step(`^request GET \/pedido with Status$`, requestGETPedidoWithStatus)
	ctx.Step(`^request DELETE \/pedido by id$`, requestDELETEPedidoById)
}

var inputs Input

type Input struct {
	pedidoID   string
	clientID   int
	statusCode int
	status     string
	firstOrder entities.Order
	body       io.ReadCloser
}

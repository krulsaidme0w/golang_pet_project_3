package httputil

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type HttpFrameworkAdapter interface {
	Error(ctx interface{}, status int, err error)
	Success(ctx interface{}, returnValue interface{})
	GetPostBody(ctx interface{}, i interface{}) error
}

type FiberFrameworkAdapter struct {
}

func NewFiberFrameworkAdapter() *FiberFrameworkAdapter {
	return &FiberFrameworkAdapter{}
}

func (adapter *FiberFrameworkAdapter) GetPostBody(ctx *fiber.Ctx, i interface{}) error {
	err := json.Unmarshal(ctx.Body(), &i)
	if err != nil {
		return err
	}
	return nil
}

func (adapter *FiberFrameworkAdapter) SetHeaders(ctx *fiber.Ctx) {
	ctx.Response().Header.Set("Content-Type", "application/json")
}

func (adapter *FiberFrameworkAdapter) SetCORSHeaders(ctx *fiber.Ctx) {
	ctx.Response().Header.Add("Access-Control-Allow-Origin", "*")
	ctx.Response().Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	ctx.Response().Header.Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

func (adapter *FiberFrameworkAdapter) Error(ctx *fiber.Ctx, status int, err_ error) {
	ctx.Context().Error(err_.Error(), status)
	if err := ctx.Status(status).SendString(err_.Error()); err != nil {
		return
	}

	adapter.SetHeaders(ctx)
	adapter.SetCORSHeaders(ctx)

	body, _ := json.Marshal(err_.Error())
	ctx.Context().SetBody(body)
}

func (adapter *FiberFrameworkAdapter) Success(ctx *fiber.Ctx, returnValue interface{}) {
	body, _ := json.Marshal(returnValue)
	ctx.Context().Success("application:json", body)
	adapter.SetHeaders(ctx)
}

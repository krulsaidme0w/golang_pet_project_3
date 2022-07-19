package httputil

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type FiberFrameworkAdapter struct {
}

func NewFiberFrameworkAdapter() *FiberFrameworkAdapter {
	return &FiberFrameworkAdapter{}
}

func (adapter *FiberFrameworkAdapter) GetPostBody(ctx *fiber.Ctx, body interface{}) error {
	err := json.Unmarshal(ctx.Body(), &body)
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
	ctx.Response().Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Response().Header.Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

func (adapter *FiberFrameworkAdapter) Error(ctx *fiber.Ctx, status int, err error) {
	ctx.Context().Error(err.Error(), status)
	if err := ctx.Status(status).SendString(err.Error()); err != nil {
		return
	}

	adapter.SetHeaders(ctx)
	adapter.SetCORSHeaders(ctx)

	body, _ := json.Marshal(err.Error())
	ctx.Context().SetBody(body)
}

func (adapter *FiberFrameworkAdapter) Success(ctx *fiber.Ctx, returnValue interface{}) {
	body, _ := json.Marshal(returnValue)
	ctx.Context().Success("application:json", body)
	adapter.SetHeaders(ctx)
}

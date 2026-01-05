package helpers

import (
	"github.com/goravel/framework/contracts/http"
)

type JsonResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Error      any    `json:"error,omitempty"`
	Errors     any    `json:"errors,omitempty"`
}

func Success(ctx http.Context, message string, data any) http.Response {
	return ctx.Response().Json(200, JsonResponse{
		StatusCode: 200,
		Message:    message,
		Data:       data,
	})
}

func Created(ctx http.Context, message string, data any) http.Response {
	return ctx.Response().Json(201, JsonResponse{
		StatusCode: 201,
		Message:    message,
		Data:       data,
	})
}

func Error(ctx http.Context, statusCode int, message string, err any) http.Response {
	return ctx.Response().Json(statusCode, JsonResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	})
}

func ValidationFailed(ctx http.Context, errors any) http.Response {
	return ctx.Response().Json(400, JsonResponse{
		StatusCode: 400,
		Message:    "Validation failed",
		Errors:     errors,
	})
}

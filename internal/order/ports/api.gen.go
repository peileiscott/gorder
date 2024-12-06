// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package ports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /customers/{customer_id}/orders)
	CreateOrder(c *gin.Context, customerID string)

	// (GET /customers/{customer_id}/orders/{order_id})
	GetOrder(c *gin.Context, customerID string, orderID string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// CreateOrder operation middleware
func (siw *ServerInterfaceWrapper) CreateOrder(c *gin.Context) {

	var err error

	// ------------- Path parameter "customer_id" -------------
	var customerID string

	err = runtime.BindStyledParameterWithOptions("simple", "customer_id", c.Param("customer_id"), &customerID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter customer_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateOrder(c, customerID)
}

// GetOrder operation middleware
func (siw *ServerInterfaceWrapper) GetOrder(c *gin.Context) {

	var err error

	// ------------- Path parameter "customer_id" -------------
	var customerID string

	err = runtime.BindStyledParameterWithOptions("simple", "customer_id", c.Param("customer_id"), &customerID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter customer_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "order_id" -------------
	var orderID string

	err = runtime.BindStyledParameterWithOptions("simple", "order_id", c.Param("order_id"), &orderID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter order_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOrder(c, customerID, orderID)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/customers/:customer_id/orders", wrapper.CreateOrder)
	router.GET(options.BaseURL+"/customers/:customer_id/orders/:order_id", wrapper.GetOrder)
}

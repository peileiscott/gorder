package ports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/order/app"
	"github.com/peileiscott/gorder/order/app/command"
	"github.com/peileiscott/gorder/order/app/query"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (s HTTPServer) PostCustomersCustomerIdOrders(c *gin.Context, customerId string) {
	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	order, err := s.app.Commands.CreateOrder.Handle(c, command.CreateOrder{
		CustomerID: req.CustomerID,
		Items:      req.Items,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"customer_id": req.CustomerID,
		"order_id":    order.OrderID,
	})
}

func (s HTTPServer) GetCustomersCustomerIdOrdersOrderId(c *gin.Context, customerID string, orderID string) {
	order, err := s.app.Queries.GetCustomerOrder.Handle(c, query.GetCustomerOrder{
		CustomerID: customerID,
		OrderID:    orderID,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}

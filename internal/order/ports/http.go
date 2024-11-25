package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/peileiscott/gorder/order/app"
	"github.com/peileiscott/gorder/order/app/query"
	"net/http"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{app: app}
}

func (s HTTPServer) PostCustomersCustomerIdOrders(c *gin.Context, customerId string) {
	//TODO implement me
	panic("implement me")
}

func (s HTTPServer) GetCustomersCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {
	order, err := s.app.Queries.GetCustomerOrder.Handle(c, query.GetCustomerOrder{
		OrderID:    "fakeID",
		CustomerID: "fakeCustomerID",
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}

package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/peileiscott/gorder/order/app"
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
	//TODO implement me
	panic("implement me")
}

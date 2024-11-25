package ports

import (
	"github.com/gin-gonic/gin"
)

type HTTPServer struct{}

func (s HTTPServer) PostCustomersCustomerIdOrders(c *gin.Context, customerId string) {
	//TODO implement me
	panic("implement me")
}

func (s HTTPServer) GetCustomersCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {
	//TODO implement me
	panic("implement me")
}

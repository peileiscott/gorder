package ports

import "github.com/gin-gonic/gin"

type HTTPServer struct{}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s HTTPServer) CreateOrder(c *gin.Context, customerID string) {
	//TODO implement me
	panic("implement me")
}

func (s HTTPServer) GetOrder(c *gin.Context, customerID string, orderID string) {
	//TODO implement me
	panic("implement me")
}

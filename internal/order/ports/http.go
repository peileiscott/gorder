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

func (s HTTPServer) CreateOrder(c *gin.Context, customerID string) {
	//TODO implement me
	panic("implement me")
}

func (s HTTPServer) GetOrder(c *gin.Context, customerID string, orderID string) {
	//TODO implement me
	panic("implement me")
}

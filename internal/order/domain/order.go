package domain

import "github.com/peileiscott/gorder/common/genproto/orderpb"

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	Items       []*orderpb.Item
	PaymentLink string
}

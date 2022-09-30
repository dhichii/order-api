package request

import (
	"order-api/src/app/order/repository/record"
)

type Request struct {
	OrderedAt    string        `json:"orderedAt"`
	CustomerName string        `json:"customerName"`
	Items        []ItemRequest `json:"items"`
}

type ItemRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (r Request) MapToRecord() record.Order {
	return record.Order{
		OrderedAt:    r.OrderedAt,
		CustomerName: r.CustomerName,
		Items:        mapItemsRequest(r.Items),
	}
}

func mapItemsRequest(request []ItemRequest) (result []record.Item) {
	for _, item := range request {
		result = append(result, mapItemRequest(item))
	}
	return
}

func mapItemRequest(item ItemRequest) record.Item {
	return record.Item{
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
}

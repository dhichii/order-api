package response

import "order-api/src/app/order/repository/record"

type Order struct {
	OrderID      int    `json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
	Items        []Item `json:"items"`
}

type Item struct {
	ItemID      int    `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func MapToBatchResponse(records []record.Order) (responses []Order) {
	for _, order := range records {
		responses = append(responses, MapToResponse(order))
	}
	return
}

func MapToResponse(record record.Order) Order {
	return Order{
		OrderID:      record.OrderID,
		CustomerName: record.CustomerName,
		OrderedAt:    record.OrderedAt,
		Items:        MapItemRecords(record.Items),
	}
}

func MapItemRecords(records []record.Item) (items []Item) {
	for _, item := range records {
		items = append(items, MapItemRecord(item))
	}
	return
}

func MapItemRecord(item record.Item) Item {
	return Item{
		ItemID:      item.ItemID,
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
}

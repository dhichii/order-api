package record

import "order-api/src/app/order/handler/response"

type Order struct {
	OrderID      int    `gorm:"primaryKey"`
	CustomerName string `gorm:"type:varchar"`
	OrderedAt    string
	Items        []Item
}

type Item struct {
	ItemID      int    `gorm:"primaryKey"`
	ItemCode    string `gorm:"type:varchar"`
	Description string
	Quantity    int
	OrderID     int
}

func MapToBatchResponse(records []Order) (responses []response.Order) {
	for _, order := range records {
		responses = append(responses, MapToResponse(order))
	}
	return
}

func MapToResponse(record Order) response.Order {
	return response.Order{
		OrderID:      record.OrderID,
		CustomerName: record.CustomerName,
		OrderedAt:    record.OrderedAt,
		Items:        MapItemRecords(record.Items),
	}
}

func MapItemRecords(records []Item) (items []response.Item) {
	for _, item := range records {
		items = append(items, MapItemRecord(item))
	}
	return
}

func MapItemRecord(item Item) response.Item {
	return response.Item{
		ItemID:      item.ItemID,
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
}

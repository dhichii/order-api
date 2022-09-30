package response

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

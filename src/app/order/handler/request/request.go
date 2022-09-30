package request

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

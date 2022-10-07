package record

type Order struct {
	OrderID      int    `gorm:"primaryKey"`
	CustomerName string `gorm:"type:varchar"`
	OrderedAt    string
	Items        []Item `gorm:"constraint:OnDelete:CASCADE"`
}

type Item struct {
	ItemID      int    `gorm:"primaryKey"`
	ItemCode    string `gorm:"type:varchar"`
	Description string
	Quantity    int
	OrderID     int
}

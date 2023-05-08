package entitys

type Product struct {
	ID          int `gorm:"prikaryKey;autoIncrement"`
	Description string
	Price       float64
	CreatorID   int
	Creator     Creator
}

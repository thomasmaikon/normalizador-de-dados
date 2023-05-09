package entitys

type Product struct {
	ID          int    `gorm:"prikaryKey;autoIncrement"`
	Description string `gorm:"unique"`
	Price       float64
	CreatorID   int
	Creator     Creator
}

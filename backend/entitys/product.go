package entitys

type Product struct {
	ID          int `gorm:"prikaryKey;autoIncrement"`
	Description string
	CreatorID   int
	Creator     Creator
}

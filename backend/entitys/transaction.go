package entitys

type Transaction struct {
	ID          int `gorm:"primaryKey"`
	Description string
	KeyFeature  string
	Signal      string
}

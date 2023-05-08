package entitys

type Afiliated struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	LeftOver  float64
	CreatorID int
	Creator   Creator
}

package entitys

type Creator struct {
	ID       int `gorm:"autoIncrement;primaryKey"`
	Name     string
	LeftOver float64
	UserID   int
	User     User
}

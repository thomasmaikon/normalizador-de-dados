package entitys

type Creator struct {
	ID       int `gorm:"autoIncrement;primaryKey"`
	Name     string
	LeftOver uint64
	UserID   int
	User     User
}

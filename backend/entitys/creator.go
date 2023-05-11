package entitys

type Creator struct {
	ID       int `gorm:"autoIncrement;primaryKey"`
	Name     string
	UserID   int
	User     User
}

package entitys

type Afiliated struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	Name      string  
	CreatorID int
	Creator   Creator
}

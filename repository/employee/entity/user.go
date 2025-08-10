package entity

type User struct {
	ID           int64 `gorm:"primaryKey"`
	Username     string
	PasswordHash string
}

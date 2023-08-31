package entities

type Users struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Address   string
	Birthdate string
	City      string
}

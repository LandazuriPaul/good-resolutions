package models

// User defines a user for the app
type User struct {
	BaseModelSoftDelete
	Email       string `gorm:"not null;unique_index:idx_email"`
	Password    *string
	UserID      *string // External user ID
	NickName    *string
	FirstName   *string
	LastName    *string
	Location    *string `gorm:"size:1048"`
	Description *string `gorm:"size:1048"`
}

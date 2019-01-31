package user

type User struct {
	UserSeq  int `gorm:"primary_key"`
	UserId   string
	Password string
	Name     string
	Company  string
	Position string
	Duty     string
	Phone    string
	Email    string
}

type UserParam struct {
	data   User
	limit  int
	offset int
}

func (User) TableName() string {
	return "info.user"
}

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
	// User
	UserId   string
	Password string
	Limit    int
	Offset   int
}

type UserPassword struct {
	UserSeq      int
	UserId       string
	Password     string
	PrevPassword string
}

func (User) TableName() string {
	return "info.user"
}

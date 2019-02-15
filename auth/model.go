package auth

type Auth struct {
	Id            string `gorm:"primary_key"`
	Grant_type    string
	Access_token  string
	Refresh_token string
}

func (Auth) TableName() string {
	return "auth_token"
}

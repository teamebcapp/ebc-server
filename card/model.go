package card

type BusinessCard struct {
	BcSeq    int `gorm:"primary_key"`
	UserId   string
	Name     string
	Company  string
	Depart   string
	Team     string
	Position string
	Duty     string
	Phone    string
	Tel      string
	Fax      string
	Address  string
	Email    string
	Priority int //`gorm:"default:'9'"`
}

type BusinessCardParam struct {
	// User
	BcSeq  int    //`schema:"bc_seq"`
	UserId string //`schema:"user_id"`
	Limit  int
	Offset int
}

func (BusinessCard) TableName() string {
	return "info.business_card"
}

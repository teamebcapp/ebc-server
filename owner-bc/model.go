package ownerbc

type OwnerBc struct {
	OwnerSeq    int `gorm:"primary_key"`
	OwnerBcSeq  int
	OwnerUserId string
	Name        string
	Company     string
	Depart      string
	Team        string
	Position    string
	Duty        string
	Phone       string
	Tel         string
	Fax         string
	Address     string
	Email       string
}

type OwnerBcParam struct {
	// User
	OwnerSeq    int
	OwnerBcSeq  int
	OwnerUserId string
	Limit       int
	Offset      int
}

func (OwnerBc) TableName() string {
	return "info.owner_bc"
}

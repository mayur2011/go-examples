package domain

type User struct {
	ID    string
	Name  string
	Email string
	Contact
	Address []Address
}

type Contact struct {
	CoutryCode string
	Number     int
}

type Address struct {
	Type  string
	Line1 string
	Line2 string
	Line3 string
	Line4 string
}

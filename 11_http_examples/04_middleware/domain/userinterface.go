package domain

type UserStore interface {
	Create(User) (string, error)
	Delete(string) error
	//RecordByID(string) (User, error)
	GetAll() ([]User, error)
}

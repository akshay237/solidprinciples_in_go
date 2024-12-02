package ocp

type User struct {
	id    int
	name  string
	phone int64
}

func (u *User) GetSender() *User {
	return &User{
		id:    1,
		name:  "Ram",
		phone: 9876543210,
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    2,
		name:  "Shyam",
		phone: 9753186420,
	}
}

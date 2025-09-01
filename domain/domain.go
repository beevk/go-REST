package domain

type UserRepo interface {
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetById(id int64) (*User, error)
	Create(user *User) (*User, error)
}

type ToDoRepo interface {
	Create(todo *ToDo) (*ToDo, error)
	Update(todo *ToDo) (*ToDo, error)
	GetByUserId(id int64) (*ToDo, error)
}

type DB struct {
	UserRepo UserRepo
	ToDoRepo ToDoRepo
}
type Domain struct {
	DB *DB
}

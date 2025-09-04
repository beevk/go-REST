package domain

type UserRepo interface {
	GetById(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) (*User, error)
}

type ToDoRepo interface {
	GetById(id int64) (*ToDo, error)
	GetByUserId(id int64) ([]*ToDo, error)
	Create(todo *ToDo) (*ToDo, error)
	Update(todo *ToDo) (*ToDo, error)
	Delete(todo *ToDo) error
}

type HasOwner interface {
	IsOwner(user *User) bool
}

type DB struct {
	UserRepo UserRepo
	ToDoRepo ToDoRepo
}
type Domain struct {
	DB *DB
}

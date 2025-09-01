package storage

import (
	"errors"

	"github.com/beevk/go-todo/domain"
	"github.com/go-pg/pg/v10"
)

type ToDoRepo struct {
	DB *pg.DB
}

func NewToDoRepo(db *pg.DB) *ToDoRepo {
	return &ToDoRepo{DB: db}
}

func (t ToDoRepo) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t ToDoRepo) Update(todo *domain.ToDo) (*domain.ToDo, error) {
	_, err := t.DB.Model(todo).Where("id = ?", todo.ID).Update()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return todo, nil
}

func (t ToDoRepo) GetByUserId(id int64) (*domain.ToDo, error) {
	todo := new(domain.ToDo)

	err := t.DB.Model(todo).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return todo, nil
}

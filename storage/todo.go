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

func (t ToDoRepo) GetByUserId(id int64) ([]*domain.ToDo, error) {
	var todos []*domain.ToDo

	err := t.DB.Model(&todos).Where("user_id = ?", id).Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return todos, nil
}

func (t ToDoRepo) GetById(id int64) (*domain.ToDo, error) {
	var todo *domain.ToDo

	err := t.DB.Model(&todo).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return todo, nil
}

func (t ToDoRepo) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t ToDoRepo) Update(todo *domain.ToDo) (*domain.ToDo, error) {
	_, err := t.DB.Model(todo).Where("id = ?", todo.ID).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t ToDoRepo) Delete(todo *domain.ToDo) error {
	_, err := t.DB.Model(todo).Delete()
	if err != nil {
		return err
	}

	return nil
}

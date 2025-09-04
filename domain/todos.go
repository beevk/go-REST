package domain

import (
	"time"
)

type ToDo struct {
	tableName struct{} `pg:"todos"`
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Completed bool     `json:"completed"`
	UserID    int64    `json:"userId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateToDoPayload struct {
	Title string `json:"title"`
}

type UpdateToDoPayload struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
}

func (c *CreateToDoPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	v.MustNotBeEmpty("title", c.Title)
	v.MustBeLongerThan("title", c.Title, 3)

	return v.HasErrors(), v.errors
}

func (u *UpdateToDoPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	if u.Title != nil && *u.Title != "" {
		v.MustBeLongerThan("title", *u.Title, 3)
	}

	return v.HasErrors(), v.errors
}

func (d *Domain) Create(p CreateToDoPayload, u *User) (*ToDo, error) {
	data := &ToDo{
		Title:     p.Title,
		Completed: false,
		UserID:    u.ID,
	}

	todo, err := d.DB.ToDoRepo.Create(data)
	if err != nil {
		// Fixme:: This error will be Custom Error
		return nil, err
	}
	return todo, nil
}

func (d *Domain) GetAll(u *User) ([]*ToDo, error) {
	todos, err := d.DB.ToDoRepo.GetByUserId(u.ID)
	if err != nil {
		return nil, err // do other stuffs like logging, etc.
	}
	return todos, nil
}

func (d *Domain) Get(id int64) (*ToDo, error) {
	data, err := d.DB.ToDoRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// You can use this hook if you want to update the UpdatedAt field automatically
// But it introduces go-pg dependency in the domain layer
//func (t *ToDo) BeforeUpdate(ctx context.Context) (context.Context, error) {
//	t.UpdatedAt = time.Now()
//	return ctx, nil
//}

func (d *Domain) Update(t *ToDo, p *UpdateToDoPayload) (*ToDo, error) {
	isUpdated := false
	if *p.Title != "" {
		t.Title = *p.Title
		isUpdated = true
	}

	if p.Completed != nil {
		t.Completed = *p.Completed
		isUpdated = true
	}

	if isUpdated {
		t.UpdatedAt = time.Now()
	}

	todo, err := d.DB.ToDoRepo.Update(t)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (d *Domain) Delete(t *ToDo) error {
	err := d.DB.ToDoRepo.Delete(t)
	if err != nil {
		return err
	}
	return nil
}

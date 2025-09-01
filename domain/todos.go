package domain

import "time"

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

func (c *CreateToDoPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	v.MustNotBeEmpty("title", c.Title)
	v.MustBeLongerThan("title", c.Title, 3)

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

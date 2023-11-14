package utils

import "errors"

func FindTodo(id string) (*Todo, error) {
	for _, t := range Todos {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, errors.New("Todo not found")
}

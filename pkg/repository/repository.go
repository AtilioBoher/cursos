package repository

import "fmt"

func NewRepository() repo {
	return repo{LastId: 0}
}

func (r *repo) StoreNewUser(name string) (int, error) {
	r.LastId++
	r.Users = append(r.Users, User{Name: name, Id: r.LastId})
	return r.LastId, nil
}

func (r *repo) GetUser(id int) (string, error) {
	for _, v := range r.Users {
		if v.Id == id {
			return v.Name, nil
		}
	}
	return "", fmt.Errorf("user with id: %v not found", id)
}
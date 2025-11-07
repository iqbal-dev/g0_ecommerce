package repo

import "errors"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Update(id int, user User) (*User, error)
	Delete(id int) (bool, error)
	FindALl() ([]*User, error)
	FindOne(id int) (*User, error)
	FindByEmailAndPassword(email string, password string) (*User, error)
}

type userRepo struct {
	userList []*User
	lastID   int
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	return repo
}

func (r *userRepo) Create(user User) (*User, error) {
	r.lastID++
	user.Id = r.lastID
	r.userList = append(r.userList, &user)
	return &user, nil
}

func (r *userRepo) Update(id int, user User) (*User, error) {
	for i := range r.userList {
		if r.userList[i].Id == id {
			user.Id = id // ensure original ID
			r.userList[i] = &user
			return &user, nil
		}
	}
	return nil, errors.New("user not found")

}

func (r *userRepo) Delete(id int) (bool, error) {
	var newList []*User
	deleted := false

	for _, u := range r.userList {
		if u.Id == id {
			deleted = true
			continue
		}
		newList = append(newList, u)
	}

	r.userList = newList
	return deleted, nil
}

func (r *userRepo) FindALl() ([]*User, error) {

	if len(r.userList) == 0 {
		return make([]*User, 0), nil
	}
	return r.userList, nil

}

func (r *userRepo) FindOne(id int) (*User, error) {
	for _, u := range r.userList {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *userRepo) FindByEmailAndPassword(email string, password string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

package memberships

import (
	"spotify_api/internal/model/memberhsips"
)

func (r *Repository) CreateUser(model memberhsips.User) error {
	return r.db.Create(model).Error
}

func (r *Repository) GetUser(username, email string, id uint) (*memberhsips.User, error) {
	user := memberhsips.User{}
	res := r.db.Where("id = ?", id).Or("username = ?", username).Or("email = ?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

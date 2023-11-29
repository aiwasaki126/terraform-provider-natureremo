package entity

// User is the user entity.
type User struct {
	id       string
	nickname string
}

// ReconstructUser is the entity constructor from data in DB.
// This constructor is assumed to be used in repository.
func ReconstructUser(id, nickname string) (*User, error) {
	return &User{
		id:       id,
		nickname: nickname,
	}, nil
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetNickname() string {
	return u.nickname
}

func (u *User) SetNickname(nickname string) error {
	u.nickname = nickname
	return nil
}

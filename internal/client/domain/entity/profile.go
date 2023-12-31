package entity

import "terraform-provider-natureremo/internal/client/domain/valueobject"

// Profile is the user profile entity.
type Profile struct {
	User
	valueobject.Preference
}

// ReconstructProfile is the entity constructor from data in DB.
// This constructor is assumed to be used in repository.
func ReconstructProfile(u User, p valueobject.Preference) (*Profile, error) {
	profile := &Profile{
		User:       u,
		Preference: p,
	}
	return profile, nil
}

func (u *Profile) GetNickname() string {
	return u.User.GetNickname()
}

func (u *Profile) UpdateNickname(v string) error {
	return u.SetNickname(v)
}

func (u *Profile) GetPreference() valueobject.Preference {
	return u.Preference
}

func (u *Profile) UpdatePreference(v valueobject.Preference) error {
	u.Preference = v
	return nil
}

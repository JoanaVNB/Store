package service

import domain "BE-JoanaVidon/user-api/domain"

func Validate(u domain.User) bool{
	if (u.Name == "" || u.CPF == "" || u.Email == "" || u.PhoneNumber == ""){
		return false
	}
	return true
}

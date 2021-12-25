package model

type User struct {
	Id 			int
	Name 		string
	Surname 	string
}

func GetAllUsers() (users []User, err error){
	users = []User{
		{1, "Name1", "Surname1"},
		{2, "Name2", "Surname2"},
		{3, "Name3", "Surname3"},
		{4, "Name4", "Surname4"},
		{5, "Name5", "Surname5"},
		{6, "Name6", "Surname6"},
		{7, "Name7", "Surname7"},
		{8, "Name8", "Surname8"},
		{9, "Name9", "Surname9"},
		{10, "Name10", "Surname10"},
	}

	return
}
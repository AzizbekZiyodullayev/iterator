package main

type userCollection struct {
	users []*User
}

func (u *userCollection) createIteraction() iterator {
	return &userIterator{
		users: u.users,
	}
}

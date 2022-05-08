package iterator

type UserIterator struct {
	User  []*User
	index int
}

func (u *UserIterator) HasNext() bool {
	if u.index < len(u.User) {
		return true
	}
	return false
}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.User[u.index]
		u.index++
		return user
	}
	return nil
}

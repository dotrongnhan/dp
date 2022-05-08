package flyweight

type Context struct {
	Id   string
	Star int
}

func NewContext(id string, star int) *Context {
	return &Context{
		Id:   id,
		Star: star,
	}
}

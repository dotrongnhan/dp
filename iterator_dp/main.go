package main

import (
	"Iterator/iterator"
	"fmt"
)

func main() {
	user1 := &iterator.User{Name: "nhan", Age: 25}
	user2 := &iterator.User{Name: "tam", Age: 17}
	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}
	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

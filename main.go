package main

import "fmt"

func main() {
	user1 := &User{
		name: "Azizbek",
		age:  25,
	}
	user2 := &User{
		name: "Kamola",
		age:  22,
	}
	user3 := &User{
		name: "Shukrona",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*User{user1, user2, user3},
	}

	iterator := userCollection.createIteraction()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}

package iterator

import (
    "fmt"
    "testing"
)

func Test_Iterator(t *testing.T) {
    user1 := &user{
        name: "a",
        age:  30,
    }
    user2 := &user{
        name: "b",
        age:  20,
    }
    userCollection := &userCollection{
        users: []*user{user1, user2},
    }
    iterator := userCollection.createIterator()
    for iterator.hasNext() {
        user := iterator.getNext()
        fmt.Printf("User is %+v\n", user)
    }
}

/*
User is &{name:a age:30}
User is &{name:b age:20}
*/
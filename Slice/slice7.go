package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriend := make([]string, len(friends))

	fmt.Println("Friend: , yourFriend:", friends, yourFriend)
	copy(yourFriend, friends)
	fmt.Println("Friend: , yourFriend:", friends, yourFriend)
	//friends[0] = "Ronald"
	fmt.Println("Friend: , yourFriend:", friends, yourFriend)
	yourFriend[0] = "Ronald"
	fmt.Println("Friend: , yourFriend:", friends, yourFriend)

	yourFriend1 := []string{}
	yourFriend1 = append(yourFriend1, friends...)
	fmt.Println(friends, yourFriend, yourFriend1)
	yourFriend1[0] = "Sonu"
	fmt.Println(friends, yourFriend1)
}

package main

import "fmt"

type PersonInfo struct {
	id      string
	name    string
	address string
}

func main() {
	map1()
}

func map1() {
	//var personDB map[string] PersonInfo
	personDB := make(map[string]PersonInfo)
	// 往这个map中插入数据
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101"}
	personDB["2"] = PersonInfo{"2", "Rose", "Room 102"}
	personDB["A"] = PersonInfo{"id_a", "Jerry", "A_Room 103"}
	personDB["B"] = PersonInfo{"id_b", "Tom", "B_Room 104"}

	person, ok := personDB["A"] // ok
	if ok {
		fmt.Println("Found person", person.name, "with id A")
	} else {
		fmt.Println("NO person with the id")
	}
}

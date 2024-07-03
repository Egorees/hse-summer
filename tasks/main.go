package main

import "fmt"

func main() {
	//testing tasks

	ht := NewHashTable(13)
	ht.Add("Egor", 19)
	ht.Add("Camila", 17)
	ht.Add("Ilya", 18)
	fmt.Println(ht.GetValue("Egor"))
	ht.Delete("Egor")
	fmt.Println(ht.GetSize())
	fmt.Println(ht.GetValue("Egor"))
	fmt.Println(ht)
}

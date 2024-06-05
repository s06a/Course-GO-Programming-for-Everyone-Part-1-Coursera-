package main

import "fmt"

func main() {
	var phoneNumber = []uint64{123456, 234567, 345678}
	var name = []string{"Jim", "Jack", "John"}
	phone := make(map[string]uint64)
	var nameInfo string
	for i := 0; i < len(name); i++ {
		phone[name[i]] = phoneNumber[i]
	}
	fmt.Println("Enter a name to lookup")
	fmt.Scanf("%s\n", &nameInfo)
	fmt.Println("phone number is", phone[nameInfo])
}
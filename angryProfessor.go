package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	total := int32(0)
	for _, att := range a {
		if att <= 0 {
			total++
		}
	}

	if total >= k {
		return "NO"
	} else {
		return "YES"
	}
}

func main() {
	// validate classroom cancel or not
	k := int32(2)              // minimum student attendance
	a := []int32{-1, -3, 4, 2} // time of student attendance, 0 as classroom start
	fmt.Println("Classroom will be cancel:")
	fmt.Println(angryProfessor(k, a))
}

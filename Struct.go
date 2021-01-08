	package main
	

	import "fmt"
	

	//defining a struct type
	//Student structure
	type student struct {
		name         string
		age          int
		depart       string
		branch       string
		college      string
	}
	

	func main() {
		// Naming fields while
		// initializing a struct
		stud1 := student{
			name:         "Deepti",
			age:          18,
			depart:       "IT",
			branch:       "B.Tech",
			college:      "MIT",
		}
		//without using fields name
		stud2 := student{"kishor", 20, "CSE", "MCA", "Pondicherry University"}
		stud3 := student{"allen", 19, "MECH", "B.Tech", "MVIT"}
		fmt.Println("student 1:", stud1)
		fmt.Println("student 2:", stud2)
		fmt.Println("student 3:", stud3)
	

	}

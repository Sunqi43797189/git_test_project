package main

import "fmt"

func switch_test()  {
	var grade string = "B"
	var marks int
	fmt.Println("请输入成绩")
	fmt.Scanln(&marks)

	switch marks {
		case 90:
			grade = "A"
		case 80:
			grade = "B"
		case 50, 60, 70:
			grade = "C"
		default:
			grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B":
		fmt.Printf("良好\n")
	case grade == "C":
		fmt.Printf("及格\n")
	case grade == "D":
		fmt.Printf("不及格")
	}

	fmt.Printf("你的成绩是 %s", grade)
}

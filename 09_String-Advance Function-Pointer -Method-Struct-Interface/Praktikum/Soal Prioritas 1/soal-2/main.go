package main

import "fmt"

type Student struct {
	name  string
	score []int
}

func (s Student) averageScore() int {
	sum := 0
	for _, score := range s.score {
		sum += score
	}
	return int(sum) / int(len(s.score))
}

func (s Student) maxScore() int {
	max := s.score[0]
	for _, score := range s.score {
		if score > max {
			max = score
		}
	}
	return max
}

func (s Student) minScore() int {
	min := s.score[0]
	for _, score := range s.score {
		if score < min {
			min = score
		}
	}
	return min
}

func main() {
	students := []Student{
		{"Rizky", []int{80}},
		{"Kobar", []int{75}},
		{"Ismail", []int{70}},
		{"Umam", []int{75}},
		{"Yopan", []int{60}},
	}
	var sum int
	for _, student := range students {
		sum += student.averageScore()
	}
	fmt.Printf("Average Score: %d\n", sum/len(students))

	var max, min int
	var maxStudent, minStudent string
	for i, student := range students {
		if i == 0 {
			max, min = student.maxScore(), student.minScore()
			maxStudent, minStudent = student.name, student.name
		} else {
			if student.maxScore() > max {
				max = student.maxScore()
				maxStudent = student.name
			}
			if student.minScore() < min {
				min = student.minScore()
				minStudent = student.name
			}
		}
	}
	fmt.Printf("Min Score of Students: %s (%d)\n", minStudent, min)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent, max)
}

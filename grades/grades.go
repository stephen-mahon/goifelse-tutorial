package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const totalDays = 192

type student struct {
	name       string
	grade      rune
	percentage float32
	attDays    int
	courses    course
}

type course struct {
	dataLogic    float32
	alogrithm    float32
	calculus     float32
	finalProject int
}

type studentList [12]student

func (s *studentList) generateMarks() {
	for i := 0; i < 12; i++ {
		name := fmt.Sprintf("Student %v", i+1)
		s[i] = student{name, ' ', 0, 82 + rand.Intn(111),
			course{50 + float32(rand.Intn(51)), 50 + float32(rand.Intn(51)), 50 + float32(rand.Intn(51)), 8 + rand.Intn(18)}}
	}
}

func printStudentDetails(s []student) {
	fmt.Printf("\nStudent Name\tPrecentage\tAttendence\tProject\tFinal Grade")
	fmt.Printf("\n-------------------------------------------------------------------")
	for i := 0; i < 12; i++ {
		fmt.Printf("\n%-10v\t%8.3v\t%6v\t%12v\t%8c", s[i].name, s[i].percentage, s[i].attDays, s[i].courses.finalProject, s[i].grade)
	}
}

func runBusinessLogic(s []student) {
	var grade rune
	for i := 0; i < 12; i++ {
		s[i].percentage, grade = calGrade(s[i])
		if grade != 'F' {
			calAttendance(s[i].attDays, &grade)
		}
		grade = calProject(s[i].courses.finalProject, grade)

		s[i].grade = grade

	}
}

func calGrade(s student) (float32, rune) {
	var percent float32
	grade := 'F'

	percent = (s.courses.alogrithm + s.courses.calculus + s.courses.dataLogic) / 3
	if percent >= 90.0 {
		grade = 'A'
	} else if percent >= 80.0 && percent < 90.0 {
		grade = 'B'
	} else if percent >= 70.0 && percent < 80.0 {
		grade = 'C'
	} else if percent >= 60.0 && percent < 70.0 {
		grade = 'D'
	} else {
		grade = 'F'
	}
	return percent, grade
}

func calAttendance(attDays int, grade *rune) {
	if float32(attDays)/totalDays < 0.75 {
		// working with the unicode values of grade instead of the char value
		*grade++
	}
}

func calProject(marks int, grade rune) rune {
	percent := float32(marks) / 25 * 100
	if percent >= 80 && grade != 'A' {
		grade--
	} else if percent < 40 && grade != 'F' {
		grade++
	}
	return grade
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var Students studentList

	Students.generateMarks()

	students := Students[:]

	runBusinessLogic(students)

	// Sort list by grade
	sort.Slice(students, func(i, j int) bool {
		return students[i].grade < students[j].grade
	})

	printStudentDetails(students)
}

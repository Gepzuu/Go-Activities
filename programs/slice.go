/*
I create an empty slice of students, add some students to it, and then perform various operations 
on the slice, such as appending new students, copying the slice, and slicing it to get a subset of students
*/

package main

import (
    "fmt"
    "slices" // The package 
)

func main() {
    // Declare an empty slice of strings to represent students
    var students []string
    fmt.Println("School", students, students == nil, len(students) == 0)

    // Initialize the students slice with a capacity of 3
    students = make([]string, 3)
    fmt.Println("Room", students, "length:", len(students), "capacity:", cap(students))

    // Add students' names to the slice
    students[0] = "Jobert"
    students[1] = "Pepito"
    students[2] = "Ghost"
    fmt.Println("added students:", students)
    fmt.Println("get student:", students[2])

    // Print the length of the students slice
    fmt.Println("length:", len(students))

    // Append more students to the slice
    students = append(students, "Gnar")
    students = append(students, "Eve", "Frankie")
    fmt.Println("added more students:", students)

    // Create a copy of the students slice
    copyOfStudents := make([]string, len(students))
    copy(copyOfStudents, students)
    fmt.Println("copy of students:", copyOfStudents)

    // Extract the last three students from the slice
    lastThree := students[len(students)-3:]
    fmt.Println("lastthree students:", lastThree)

    // Extract the first five students from the slice
    firstFive := students[:5]
    fmt.Println("first five students:", firstFive)

    // Extract all students using slice syntax
    allStudents := students[:]
    fmt.Println("all students:", allStudents)

    // Create a new slice representing a new class
    newClass := []string{"Khazix", "Hecarim", "Ivern"}
    fmt.Println("new class:", newClass)

    // Compare the new class with the students slice
    if slices.Equal(newClass, students) {
        fmt.Println("new class is equal to students")
    }

    // Create a 2D slice to represent grades
    grades := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        grades[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            grades[i][j] = i + j
        }
    }
    fmt.Println("Grades:", grades)
}

/*

# Output

School [] true true
Room [  ] length: 3 capacity: 3
added students: [Jobert Pepito Ghost]
get student: Ghost
length: 3
added more students: [Jobert Pepito Ghost Gnar Eve Frankie]
copy of students: [Jobert Pepito Ghost Gnar Eve Frankie]
lastthree students: [Gnar Eve Frankie]
first five students: [Jobert Pepito Ghost Gnar Eve]
all students: [Jobert Pepito Ghost Gnar Eve Frankie]
new class: [Khazix Hecarim Ivern]
Grades: [[0] [1 2] [2 3 4]]


*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Student struct represents the structure of student information
type Student struct {
	Name string
	Id   string
	Phno int
}

// Subject struct represents the structure of subject information
type Subject struct {
	Name  string
	Id    string
	Marks int
}

type Marks struct {
	StudentID string
	SubjectID string
	Marks     float64
}

func main() {
	// Read students from file
	students, err := readStudents("students.csv")
	if err != nil {
		fmt.Println("Error reading student file:", err)
		return
	}

	// Read subjects from file
	subjects, err := readSubjects("subjects.csv")
	if err != nil {
		fmt.Println("Error reading subjects file:", err)
		return
	}

	// Read marks from file
	marks, err := readMarks("marks.csv")
	if err != nil {
		fmt.Println("Error reading marks file:", err)
		return
	}

	// Print student information
	fmt.Println("Students:")
	for _, student := range students {
		fmt.Printf("Name: %s, Id: %s, Phone: %d\n", student.Name, student.Id, student.Phno)
	}

	// Print subject information
	fmt.Println("\nSubjects:")
	for _, subject := range subjects {
		fmt.Printf("Name: %s, Id: %s, Marks: %d\n", subject.Name, subject.Id, subject.Marks)
	}

	// Print marks information
	fmt.Println("\nMarks:")
	for _, mark := range marks {
		fmt.Printf("StudentID: %s, SubjectID: %s, Marks:  %.2f\n", mark.StudentID, mark.SubjectID, mark.Marks)
	}

	//  b. Calculate and print class averages
	classAverages := calculateClassAverages(subjects, marks)
	fmt.Println("\nClass Averages:")
	for subjectID, average := range classAverages {
		fmt.Printf("SubjectID: %s, Average: %.2f\n", subjectID, average)
	}

	// b.Calculate and print total marks
	totalMarks := calculateTotalMarks(marks)
	fmt.Printf("\nTotal Marks: %.2f\n", totalMarks)

	//  c .Print total marks of each student in alphabetical order
	printTotalMarksInAlphabeticalOrder(students, marks)
    // d.to print students and their ids in order of their marks
	printStudentsInOrderOfMarks(students, marks)
	//e . given student id print their info ,marks,total and class average for each subject and total marks.
	// Provide the student ID for which you want to print information
	studentIDToPrint := "S001"
	printStudentInfo(studentIDToPrint, students, marks, subjects, classAverages)


}

// readStudents reads student information from a CSV file
func readStudents(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var students []Student
	for _, line := range lines {
		phno, _ := strconv.Atoi(line[2]) // Convert the third column to an integer (phone number)
		students = append(students, Student{
			Name: line[0],
			Id:   line[1],
			Phno: phno,
		})
	}

	return students, nil
}

// readSubjects reads subject information from a CSV file
func readSubjects(filename string) ([]Subject, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var subjects []Subject
	for _, line := range lines {
		marks, _ := strconv.Atoi(line[2]) // Convert the third column to an integer (marks)
		subjects = append(subjects, Subject{
			Name:  line[0],
			Id:    line[1],
			Marks: marks,
		})
	}

	return subjects, nil
}

func readMarks(filename string) ([]Marks, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var marks []Marks
	for _, line := range lines {
		// Assuming the third column contains marks
		marksVal, _ := strconv.ParseFloat(line[2], 64)
		marks = append(marks, Marks{
			StudentID: line[0],
			SubjectID: line[1],
			Marks:     marksVal,
		})
	}

	return marks, nil
}

func calculateClassAverages(subjects []Subject, marks []Marks) map[string]float64 {
	// Initialize a map to store cumulative marks and count for each subject
	subjectTotals := make(map[string]float64)
	subjectCounts := make(map[string]int)

	// Iterate through marks to calculate cumulative marks and count for each subject
	for _, mark := range marks {
		subjectID := mark.SubjectID
		subjectTotals[subjectID] += mark.Marks
		subjectCounts[subjectID]++
	}

	// Calculate class averages for each subject
	classAverages := make(map[string]float64)
	for _, subject := range subjects {
		subjectID := subject.Id
		total := subjectTotals[subjectID]
		count := subjectCounts[subjectID]

		// Avoid division by zero
		if count > 0 {
			classAverages[subjectID] = total / float64(count)
		} else {
			classAverages[subjectID] = 0.0
		}
	}

	return classAverages
}

// Calculate total marks for all students
func calculateTotalMarks(marks []Marks) float64 {
	var totalMarks float64

	// Iterate through marks to calculate total marks
	for _, mark := range marks {
		totalMarks += mark.Marks
	}

	return totalMarks
}

// Print total marks of each student in alphabetical order
func printTotalMarksInAlphabeticalOrder(students []Student, marks []Marks) {
	// Create a map to store total marks for each student
	totalMarksMap := make(map[string]float64)

	// Iterate through marks to calculate total marks for each student
	for _, mark := range marks {
		studentID := mark.StudentID
		totalMarksMap[studentID] += mark.Marks
	}

	// Create a slice to store sorted student names
	var studentNames []string
	for _, student := range students {
		// Exclude the header row ("Name") from the processing
		if student.Name != "Name" {
			studentNames = append(studentNames, student.Name)
		}
	}

	// Sort the student names in alphabetical order
	sort.Strings(studentNames)

	// Print total marks for each student in alphabetical order
	fmt.Println("\nTotal Marks of Each Student in Alphabetical Order:")
	for _, studentName := range studentNames {
		studentID := getStudentIDByName(students, studentName)
		totalMarks := totalMarksMap[studentID]
		fmt.Printf("Name: %s, Total Marks: %.2f\n", studentName, totalMarks)
	}
}

func getStudentIDByName(students []Student, name string) string {
	for _, student := range students {
		if student.Name == name {
			return student.Id
		}
	}
	return ""
}

// Print students and their ids in the order of their marks
func printStudentsInOrderOfMarks(students []Student, marks []Marks) {
	// Create a map to store total marks for each student
	totalMarksMap := make(map[string]float64)

	// Iterate through marks to calculate total marks for each student
	for _, mark := range marks {
		studentID := mark.StudentID
		totalMarksMap[studentID] += mark.Marks
	}

	// Create a slice to store student data (name, id, total marks)
	var studentData []struct {
		Name       string
		ID         string
		TotalMarks float64
	}

	// Populate the studentData slice
	for _, student := range students {
		// Exclude the header row ("Name") from the processing
		if student.Name != "Name" {
			studentID := student.Id
			totalMarks := totalMarksMap[studentID]
			studentData = append(studentData, struct {
				Name       string
				ID         string
				TotalMarks float64
			}{Name: student.Name, ID: studentID, TotalMarks: totalMarks})
		}
	}

	// Sort students based on their total marks in descending order
	sort.Slice(studentData, func(i, j int) bool {
		return studentData[i].TotalMarks > studentData[j].TotalMarks
	})

	// Print students and their ids in the order of their marks
	fmt.Println("\nStudents and Their IDs in the Order of Their Marks:")
	for _, student := range studentData {
		fmt.Printf("Name: %s, ID: %s, Total Marks: %.2f\n", student.Name, student.ID, student.TotalMarks)
	}
}


// GIVEN STUDENT ID PRINT INFO:

// Print student info, marks, total, and class average for each subject, and total marks
func printStudentInfo(studentID string, students []Student, marks []Marks, subjects []Subject, classAverages map[string]float64) {
	// Find the student by ID
	var student Student
	for _, s := range students {
		if s.Id == studentID {
			student = s
			break
		}
	}

	if student.Id == "" {
		fmt.Println("Student not found.")
		return
	}

	// Print student information
	fmt.Printf("\nStudent Info for ID %s:\n", studentID)
	fmt.Printf("Name: %s, ID: %s, Phone: %d\n", student.Name, student.Id, student.Phno)

	// Print marks information for the student
	fmt.Println("\nMarks:")
	for _, mark := range marks {
		if mark.StudentID == studentID {
			subjectID := mark.SubjectID
			subjectName := getSubjectNameByID(subjectID, subjects)
			fmt.Printf("Subject: %s, Marks: %.2f\n", subjectName, mark.Marks)
		}
	}

	// Print total marks for the student
	totalMarks := calculateTotalMarksForStudent(studentID, marks)
	fmt.Printf("\nTotal Marks for ID %s: %.2f\n", studentID, totalMarks)

	// Print class average for each subject
	fmt.Println("\nClass Averages:")
	for subjectID, average := range classAverages {
		subjectName := getSubjectNameByID(subjectID, subjects)
		fmt.Printf("Subject: %s, Class Average: %.2f\n", subjectName, average)
	}
}

// getSubjectNameByID returns the subject name given a subject ID
func getSubjectNameByID(subjectID string, subjects []Subject) string {
	for _, subject := range subjects {
		if subject.Id == subjectID {
			return subject.Name
		}
	}
	return ""
}

// calculateTotalMarksForStudent calculates the total marks for a given student ID
func calculateTotalMarksForStudent(studentID string, marks []Marks) float64 {
	var totalMarks float64
	for _, mark := range marks {
		if mark.StudentID == studentID {
			totalMarks += mark.Marks
		}
	}
	return totalMarks
}


package main

import (
	"fmt"
	"time"
	"math"
	"encoding/json"
	"io/ioutil"
	"sort"
	"reflect"
	"testing"
)

//I coudn't figure out why Tasks 4 and 5 don't have outputs

// Task 1

func factorial(n int, resultChan chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	resultChan <- result
}

// Task 2

// Define the Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define the Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Define the Circle struct
type Circle struct {
	Radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that takes a Shape interface and calculates area and perimeter
func printShapeInfo(s Shape) {
	fmt.Printf("Shape Info:\n")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println()
}

// Task 3

// Custom error type for FileNotFound
type FileNotFoundError struct {
	FileName string
}

// Implement the error interface for FileNotFound
func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found: %s", e.FileName)
}

// Function that simulates an operation that may return FileNotFound error
func readFile(fileName string) ([]byte, error) {
	// Simulating an operation that may fail
	// In this example, let's assume the file is not found
	return nil, &FileNotFoundError{FileName: fileName}
}

// Task 4

// Book represents the structure of a book
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// ReadJSONFile reads a JSON file and returns the content as a byte slice
func ReadJSONFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// ParseJSON parses JSON data into a slice of Book
func ParseJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}

// SortBooksByPages sorts a slice of Book by the number of pages
func SortBooksByPages(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Pages < books[j].Pages
	})
}

// UpdatePageCount updates the page count of the first book in the slice
func UpdatePageCount(books []Book, newPageCount int) {
	if len(books) > 0 {
		books[0].Pages = newPageCount
	}
}

// WriteJSONFile writes a slice of Book to a new JSON file
func WriteJSONFile(filename string, books []Book) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

//Task 5

func TestReadJSONFile(t *testing.T) {
	expected := []byte(`{"title": "Book1", "author": "Author1", "pages": 200}`)
	err := ioutil.WriteFile("test.json", expected, 0644)
	if err != nil {
		t.Fatalf("Error writing test JSON file: %v", err)
	}

	fileData, err := ReadJSONFile("test.json")
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	if !reflect.DeepEqual(fileData, expected) {
		t.Errorf("ReadJSONFile failed. Expected: %s, Got: %s", expected, fileData)
	}
}

func TestParseJSON(t *testing.T) {
	expected := []Book{{Title: "Book1", Author: "Author1", Pages: 200}}
	data := []byte(`[{"title": "Book1", "author": "Author1", "pages": 200}]`)
	books, err := ParseJSON(data)
	if err != nil {
		t.Fatalf("Error parsing JSON data: %v", err)
	}

	if !reflect.DeepEqual(books, expected) {
		t.Errorf("ParseJSON failed. Expected: %v, Got: %v", expected, books)
	}
}

func TestSortBooksByPages(t *testing.T) {
	books := []Book{
		{Title: "Book1", Author: "Author1", Pages: 200},
		{Title: "Book2", Author: "Author2", Pages: 150},
		{Title: "Book3", Author: "Author3", Pages: 300},
	}

	SortBooksByPages(books)

	expected := []Book{
		{Title: "Book2", Author: "Author2", Pages: 150},
		{Title: "Book1", Author: "Author1", Pages: 200},
		{Title: "Book3", Author: "Author3", Pages: 300},
	}

	if !reflect.DeepEqual(books, expected) {
		t.Errorf("SortBooksByPages failed. Expected: %v, Got: %v", expected, books)
	}
}

func TestUpdatePageCount(t *testing.T) {
	books := []Book{{Title: "Book1", Author: "Author1", Pages: 200}}
	newPageCount := 250

	UpdatePageCount(books, newPageCount)

	if books[0].Pages != newPageCount {
		t.Errorf("UpdatePageCount failed. Expected: %d, Got: %d", newPageCount, books[0].Pages)
	}
}

func TestWriteJSONFile(t *testing.T) {
	expectedData := []Book{
		{Title: "Book1", Author: "Author1", Pages: 200},
		{Title: "Book2", Author: "Author2", Pages: 150},
		{Title: "Book3", Author: "Author3", Pages: 300},
	}

	err := WriteJSONFile("test_output.json", expectedData)
	if err != nil {
		t.Fatalf("Error writing test JSON file: %v", err)
	}

	fileData, err := ioutil.ReadFile("test_output.json")
	if err != nil {
		t.Fatalf("Error reading test JSON file: %v", err)
	}

	if !reflect.DeepEqual(fileData, expectedData) {
		t.Errorf("WriteJSONFile failed. Expected: %v, Got: %v", expectedData, fileData)
	}
}

func main() {
	// Task 1

	// Take input from the user
	var number int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Nuh uh... that's not a positive integer, please enter one:'")
		return
	}

	// Create a channel to collect results from goroutines
	resultChan := make(chan int)

	// Measure the starting time
	startTime := time.Now()

	// Launch multiple goroutines to calculate factorial
	for i := 1; i <= number; i++ {
		go factorial(i, resultChan)
	}

	// Collect results from goroutines
	totalResult := 1
	for i := 1; i <= number; i++ {
		totalResult *= <-resultChan
	}

	// Measure the ending time
	endTime := time.Now()

	// Display the result and execution time
	fmt.Printf("Factorial of %d is: %d\n", number, totalResult)
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))

	// Task 2

	// Get user input for Rectangle
	var rectWidth, rectHeight float64
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&rectWidth)
	fmt.Print("Enter the height of the rectangle: ")
	fmt.Scan(&rectHeight)

	// Create a Rectangle with user-provided parameters
	rect := Rectangle{Width: rectWidth, Height: rectHeight}

	// Get user input for Circle
	var circleRadius float64
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&circleRadius)

	// Create a Circle with user-provided parameters
	circle := Circle{Radius: circleRadius}

	// Calculate and print information for the Rectangle
	fmt.Println("Rectangle:")
	printShapeInfo(rect)

	// Calculate and print information for the Circle
	fmt.Println("Circle:")
	printShapeInfo(circle)

	// Task 3

	// Example usage of the readFile function
	var fileToRead string
	fmt.Print("Enter the name of the file to read: ")
	fmt.Scan(&fileToRead)

	content, err := readFile(fileToRead)

	// Check if the error is of type FileNotFoundError
	if err != nil {
		if fileNotFoundErr, ok := err.(*FileNotFoundError); ok {
			// Custom error handling for FileNotFound
			fmt.Println("Custom Error Handling:")
			fmt.Println(fileNotFoundErr.Error())
		} else {
			// Handle other types of errors
			fmt.Println("Unexpected Error:", err)
		}
		return
	}

	// Continue processing if no error occurred
	fmt.Printf("File content: %s\n", string(content))

	//Task 4

	// Sample JSON data (replace this with actual data)
	jsonData := `[{"title": "Book1", "author": "Author1", "pages": 200},
				  {"title": "Book2", "author": "Author2", "pages": 150},
				  {"title": "Book3", "author": "Author3", "pages": 300}]`

	// Write the sample data to a JSON file
	err = ioutil.WriteFile("input.json", []byte(jsonData), 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	// Read the JSON file
	fileData, err := ReadJSONFile("input.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse the data
	books, err := ParseJSON(fileData)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}

	// Perform operations on the data (e.g., sorting by pages)
	SortBooksByPages(books)

	// Print the sorted data
	fmt.Println("Sorted Books by Pages:")
	for _, book := range books {
		fmt.Printf("%s by %s - %d pages\n", book.Title, book.Author, book.Pages)
	}

	// Modify the data (e.g., update page count)
	UpdatePageCount(books, 250)

	// Print the modified data
	fmt.Println("\nModified Books:")
	for _, book := range books {
		fmt.Printf("%s by %s - %d pages\n", book.Title, book.Author, book.Pages)
	}

	// Write the modified data to a new JSON file
	err = WriteJSONFile("output.json", books)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("\nData written to output.json")
}
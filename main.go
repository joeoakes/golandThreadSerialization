package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Define the Student struct
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

func main() {

	student1 := Student{Name: "Alice", Age: 17, Grade: "11th"}
	student2 := Student{Name: "Johnny", Age: 18, Grade: "12th"}

	fmt.Println("Starting the function...")
	start := time.Now() // Record the start time
	fmt.Println(serializeStudent(student1))
	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Time elapsed: %v\n", elapsed.Microseconds())
	fmt.Println("Function finished!")
	fmt.Println(serializeStudent(student2))
	elapsed = time.Since(start) // Calculate the elapsed time
	fmt.Printf("Time elapsed: %v\n", elapsed.Microseconds())
	fmt.Println("Function 2 finished!")

	var wg sync.WaitGroup
	wg.Add(2)
	start = time.Now() // Record the start time
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		fmt.Println(serializeStudent(student1))
	}()

	go func() {
		defer wg.Done()
		fmt.Println(serializeStudent(student2))
	}()
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed.Microseconds())
	fmt.Println("All goroutines completed!")
}

func serializeStudent(student Student) string {
	studentJSON, _ := json.Marshal(student)
	return string(studentJSON)
}

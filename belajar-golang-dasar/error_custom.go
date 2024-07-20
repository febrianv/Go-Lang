package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "Edo" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := saveData("", nil)

	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not found error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown error:",err.Error())
		// }
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
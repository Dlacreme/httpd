package werror

import "fmt"

// Type is error type Enum
type Type int

// Error used in YANA
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// New build a new werror
func New(status int, message string) *Error {
	fmt.Printf("[%d] %s\n", status, message)
	return &Error{status, message}
}

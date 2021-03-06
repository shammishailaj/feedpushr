package common

import "fmt"

// ArrayFlags contains an array of command flags
type ArrayFlags []string

// String return the string value of a flag array
func (i *ArrayFlags) String() string {
	return fmt.Sprint(*i)
}

// Set is used to add a value to the flag array
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

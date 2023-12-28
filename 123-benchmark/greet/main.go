package greet

import "fmt"

func Greet(s string) string {
	return fmt.Sprint("Hello there ", s)
}

// Best practice is to cover with tests/examples/benchmarking
// around 70-80% of your code : use go test -cover

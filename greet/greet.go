package greet

import "fmt"

func Greet(req string) string {
	return fmt.Sprintf("Hello %s", req)
}

package greet

import "fmt"

//HelloHuman return a `Hello %s`
func HelloHuman(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

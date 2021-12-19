// Package twofer implements the logic for greeting someone
package twofer

import "fmt"

// ShareWith takes the name of the user as input and returns the greeting message
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

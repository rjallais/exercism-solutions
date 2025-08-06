// Package twofer implements a simple way to generate a two-fer phrase.
// The idea is that if you have an extra cookie (or any treat), you "share" it.
// If you know the person's name, you say "One for <name>, one for me.".
// Otherwise, you say "One for you, one for me.".
package twofer

import "fmt"

// ShareWith returns a two-fer phrase using the provided name.
// If a non-empty name is provided, it returns "One for <name>, one for me.".
// Otherwise, it returns "One for you, one for me.".
func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
	return "One for you, one for me."
}

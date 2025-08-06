package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	msg := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
    return msg
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*", numStarsPerLine)
	msg   := stars + "\n" + welcomeMsg + "\n" + stars
    return msg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
    msg = strings.TrimSpace(msg)
    return msg
}

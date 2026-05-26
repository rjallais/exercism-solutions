// leap package defines a function to check if an year is a leap year or not
package leap

// IsLeapYear return true if the received year is a leap year and false otherwise
func IsLeapYear(year int) bool {
	switch {
    case year % 400 == 0:
        return true
    case year % 100 == 0:
        return false
    case year % 4 == 0:
        return true
    default:
        return false
    }
}

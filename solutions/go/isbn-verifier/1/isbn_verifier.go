package isbnverifier

import (
    "strings"
)

func IsValidISBN(isbn string) bool {
	trimmed := strings.ReplaceAll(isbn, "-", "")
    if len(trimmed) != 10 {
        return false
    }

    var sum, d int
    for i, c := range trimmed {
        if c == 'X' {
            if i != 9 {
                return false
            }
            d = 10
        } else {
            d = int(c - '0')
        }
        sum += (10-i)*d
    }

    return sum%11 == 0
}

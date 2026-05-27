package luhn

import (
    "strings"
    "strconv"
)
    

func Valid(id string) bool {
	sum := 0
    offset := 0
    trimmed := strings.ReplaceAll(id, " ", "")
    if len(trimmed) < 2 {
        return false
    }
    if len(trimmed)%2 == 1 {
		offset++
	}
    
    for i, c := range trimmed {
        if c < '0' || c > '9' {
            return false
        }
        digit, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
		if (i+offset)%2 == 0 {
			double := digit * 2
			if double > 9 {
                sum += double - 9
            } else {
                sum += double
            }
        } else {
                sum += digit
        }
	}
    return sum%10 == 0
}

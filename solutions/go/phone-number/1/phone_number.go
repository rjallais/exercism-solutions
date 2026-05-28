package phonenumber

import (
    "fmt"
)

func Number(phoneNumber string) (string, error) {
	trimmed := ""
	for _, c := range phoneNumber {
        if c < '0' || c > '9' {
            continue
        }
        trimmed = trimmed + string(c)
    }

    if len(trimmed) == 11 && trimmed[0] == '1' {
        trimmed = trimmed[1:]
    }

	if len(trimmed) != 10 || trimmed[0] < '2' || trimmed[3] < '2' {
        return "", fmt.Errorf("invalid phone number")
    }
    
    return trimmed, nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
    num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("(%v) %v-%v", num[:3], num[3:6], num[6:]), nil
}

package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
    for val, list := range in {
        for _, letter := range list {
        	result[strings.ToLower(letter)] = val
        }        
    }

    return result
}

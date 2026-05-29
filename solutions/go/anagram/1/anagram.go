package anagram

import "strings"

func matches(base map[rune]int, s string) bool {
    freq := make(map[rune]int, len(base))
    for k, v := range base {
        freq[k] = v
    }

    for _, r := range s {
        n, ok := freq[r]; if !ok { return false }
        if n == 1 {
            delete(freq, r)
        } else {
            freq[r]--
        }
    }

    return len(freq) == 0
}

func Detect(subject string, candidates []string) []string {
    s := strings.ToUpper(subject)

    subjFreq := make(map[rune]int)
    for _, r := range s {
        subjFreq[r]++
    }

    anagrams := make([]string, 0, len(candidates))
    for _, cand := range candidates {
        c := strings.ToUpper(cand)
        if len(c) != len(s) || c == s {
            continue
        }
        if matches(subjFreq, c) {
            anagrams = append(anagrams, cand)
        }
    }

    return anagrams
}

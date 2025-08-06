package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	var re = regexp.MustCompile(`^\[TRC]|^\[DBG]|^\[INF]|^\[WRN]|^\[ERR]|^\[FTL]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	var re = regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var (
		re    = regexp.MustCompile(`.*".*(?i:password).*".*`)
		count = 0
	)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	var re = regexp.MustCompile(`end-of-line[0-9]+`)
	return strings.Join(re.Split(text, -1), "")
}

func TagWithUserName(lines []string) []string {
	var re = regexp.MustCompile(`User +(\S+)`)
	for i, line := range lines {
		if re.FindString(line) != "" {
			lines[i] = strings.Join([]string{"[USR]", re.FindStringSubmatch(line)[1], line}, " ")
		}
	}
	return lines
}

package parsinglogfiles

import (
	"fmt"
	"regexp"
	"slices"
)

func IsValidLine(text string) bool {
	valid := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}

	re := regexp.MustCompile(`^\[[A-Z]{3}\]`)
	s := re.FindString(text)

	return slices.Contains(valid, s)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~\-=*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int

	re := regexp.MustCompile(`".*[pP][aA][sS][sS][wW][oOS][rR][dD].*"`)
	for _, line := range lines {
		if re.FindString(line) != "" {
			count += 1
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var out []string

	re := regexp.MustCompile(`User +([a-zA-Z0-9]+) `)
	for _, line := range lines {
		sl := re.FindStringSubmatch(line)
		if sl == nil {
			out = append(out, line)
		} else {
			out = append(out, fmt.Sprintf("[USR] %s %s", sl[1], line))
		}
	}

	return out
}

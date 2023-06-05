package algorithm

import "strings"

func ReplaceSpace(s string) string {
	var result []rune
	const badSpace = '\u0020'
	for _, r := range s {
		if r == badSpace {
			result = append(result, '\u00A0')
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func TrimAll(str string) string {
	newStr := strings.ReplaceAll(str, "\n", "")
	newStr = strings.ReplaceAll(newStr, "\t", "")
	newStr = strings.ReplaceAll(newStr, " ", "")
	newStr = strings.ReplaceAll(newStr, "\u00a0", " ")

	return newStr

}

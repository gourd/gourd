package compile

import (
	"regexp"
	"strings"
)

// SubfixFn return a normalizer which takes a string, subfix it and add .go to the end
func SubfixFn(subfix string) func(string) string {
	return func(str string) string {
		r1 := regexp.MustCompile("[A-Z]+")
		r2 := regexp.MustCompile("^\\_")
		return strings.ToLower(r2.ReplaceAllString(r1.ReplaceAllString(str, "_$0"), "")) + "_" + subfix + ".go"
	}
}

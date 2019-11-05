package twitter

import (
	//"fmt"
	 "unicode/utf8"
	 "regexp"
	 "strings"
)
const min int= 1
const max int = 15
const reg string = "^[a-zA-Z0-9_]*$"
var legalCharRegexp = regexp.MustCompile(reg)

func isLongEnough(username string) bool{
	var nb = utf8.RuneCountInString(username)

	if nb > min{
		return true
	}
	return false
}

func isShortEnough(username string) bool{
	var nb = utf8.RuneCountInString(username)

	if nb < max{
		return true
	}
	return false
}

func onlyContainsLegalChars(username string) bool{
	return legalCharRegexp.MatchString(username)
}

func containsNoIllegalPattern(username string) bool{
	return !strings.Contains(strings.ToLower(username), "twitter")
}

func Validate(username string) bool{
	return isLongEnough(username) &&
	 containsNoIllegalPattern(username) &&
	  onlyContainsLegalChars(username) &&
	   isShortEnough(username) &&
	    isLongEnough(username)
}
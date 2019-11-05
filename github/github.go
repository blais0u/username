package github

import (
	//"fmt"
	 "unicode/utf8"
	 "regexp"
	 "strings"
)
const min int= 1
const max int = 39
const reg string = "^[-0-9a-zA-Z]*$"
var legalCharRegexp = regexp.MustCompile(reg)
//prefix et suffix illegaux : -
// patern illegal -
// min 1 
// max 39
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
	return !strings.HasPrefix(username, "-") &&
	 !strings.HasSuffix(username, "-")  &&
	 !strings.Contains(username, "_") &&
	  !strings.Contains(username, "--")

}

func Validate(username string) bool{
	return isLongEnough(username) &&
	 containsNoIllegalPattern(username) &&
	  onlyContainsLegalChars(username) &&
	   isShortEnough(username) &&
	    isLongEnough(username)
}
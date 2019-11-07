package github

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/blais0u/username"
)

const min int = 1
const max int = 39
const reg string = "^[-0-9a-zA-Z]*$"

var legalCharRegexp = regexp.MustCompile(reg)

type GitHub struct {
}

func init() {
	var git = GitHub{}
	username.SocialNetWorks = append(username.SocialNetWorks, &git)
}

//prefix et suffix illegaux : -go get _u
// patern illegal -
// min 1
// max 39
func isLongEnough(pseudo string) bool {
	var nb = utf8.RuneCountInString(pseudo)

	if nb > min {
		return true
	}
	return false
}

func isShortEnough(pseudo string) bool {
	var nb = utf8.RuneCountInString(pseudo)

	if nb < max {
		return true
	}
	return false
}

func onlyContainsLegalChars(pseudo string) bool {
	return legalCharRegexp.MatchString(pseudo)
}

func containsNoIllegalPattern(pseudo string) bool {
	return !strings.HasPrefix(pseudo, "-") &&
		!strings.HasSuffix(pseudo, "-") &&
		!strings.Contains(pseudo, "_") &&
		!strings.Contains(pseudo, "--")

}

func (*GitHub) Validate(pseudo string) bool {
	return isLongEnough(pseudo) &&
		containsNoIllegalPattern(pseudo) &&
		onlyContainsLegalChars(pseudo) &&
		isShortEnough(pseudo) &&
		isLongEnough(pseudo)
}
func (*GitHub) IsAvailable(pseudo string) (ok bool, err error) {
	resp, err := http.Get("https://github.com/" + pseudo)
	if err != nil {
		//read response status
		errAvail := username.ErrAvailibility{
			Cause:         err,
			SocialWebsite: "GitHub",
		}
		return false, &errAvail
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return true, err
	}

	return false, err
}

package twitter

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/blais0u/username"
)

const min int = 1
const max int = 15
const reg string = "^[a-zA-Z0-9_]*$"

type Twitter struct {
}

var legalCharRegexp = regexp.MustCompile(reg)

func init() {
	var tw = Twitter{}
	username.SocialNetWorks = append(username.SocialNetWorks, &tw)
}
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
	return !strings.Contains(strings.ToLower(pseudo), "twitter")
}

func (*Twitter) Validate(pseudo string) bool {
	return isLongEnough(pseudo) &&
		containsNoIllegalPattern(pseudo) &&
		onlyContainsLegalChars(pseudo) &&
		isShortEnough(pseudo) &&
		isLongEnough(pseudo)
}

func (*Twitter) IsAvailable(pseudo string) (ok bool, err error) {
	resp, err := http.Get("https://twitter.com/" + pseudo)
	if err != nil {
		//read response status
		errAvail := username.ErrAvailibility{
			Cause:         err,
			SocialWebsite: "Twitter",
		}
		return false, &errAvail
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return true, err
	}

	return false, err
}

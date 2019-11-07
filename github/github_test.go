package github

import "testing"

var git = GitHub{}

func TestIsLongEnough(t *testing.T) {
	if git.Validate("dssqdsqdsqdsqdzaezaezaezaezaezasqdsqdsqd") {
		t.Error("TooLong is on error")
	}
}

func TestIsShortEnough(t *testing.T) {
	if git.Validate("a") {
		t.Error("Tooshort is on error")

	}
}

func TestContainsLegalChars(t *testing.T) {
	if git.Validate("éSalutà") {
		t.Error("IllegalChar is on error")

	}
}

func TestContainsLegalCharsWithUnderScore(t *testing.T) {
	if git.Validate("-Salut_") {
		t.Error("IllegalChar is on error")

	}
}
func TestContainsNoIllegalPattern(t *testing.T) {
	if git.Validate("-SalutTwittEr") {
		t.Error("Illegalpattern is on error")
	}
}

func TestValidate(t *testing.T) {
	if !git.Validate("Bonjour") {
		t.Error("Validate is on error")
	}
}

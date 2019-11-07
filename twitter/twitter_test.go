package twitter

import "testing"

var tw = Twitter{}

func TestIsLongEnough(t *testing.T) {

	if tw.Validate("dssqdsqdsqdsqdsqdsqdsqd") {
		t.Error("TooLong is on error")
	}
}

func TestIsShortEnough(t *testing.T) {
	if tw.Validate("a") {
		t.Error("Tooshort is on error")

	}
}

func TestContainsLegalChars(t *testing.T) {
	if tw.Validate("éSalutà") {
		t.Error("IllegalChar is on error")

	}
}
func TestContainsNoIllegalPattern(t *testing.T) {
	if tw.Validate("SalutTwittEr") {
		t.Error("Illegalpattern is on error")

	}
}

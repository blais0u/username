package twitter
import "testing"
func TestIsLongEnough(t *testing.T){
	if (Validate("dssqdsqdsqdsqdsqdsqdsqd")){
		t.Error("TooLong is on error")
	}
}

func TestIsShortEnough(t *testing.T){
	if (Validate("a")){
		t.Error("Tooshort is on error")

	}
}

func TestContainsLegalChars(t *testing.T){
	if (Validate("éSalutà")){
		t.Error("IllegalChar is on error")

	}
}
func  TestContainsNoIllegalPattern(t *testing.T){
	if (Validate("SalutTwittEr")){
		t.Error("Illegalpattern is on error")

	}
}
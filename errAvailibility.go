package username

import "fmt"

//wrapper et error en interface
type ErrAvailibility struct {
	SocialWebsite string
	Cause         error
	PseudoTested  string
}

func (e *ErrAvailibility) Unwrap() error {
	return e.Cause
}
func (e *ErrAvailibility) Error() string {
	return fmt.Sprintf("Error on" + e.SocialWebsite + " due to" + e.Cause.Error())
}

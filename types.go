package username

// tyoe interface validateyr a la racine
// type interface

type IsAvailabler interface {
	IsAvailable(username string) (ok bool, err error)
}

type Validator interface {
	Validate(username string) bool
}

type SocialNetwork interface {
	Validator
	IsAvailabler
}

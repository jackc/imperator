package validate

type Validator struct {
	e Errors
}

func (v *Validator) Add(err Error) {
	if v.e == nil {
		v.e = make(Errors)
	}

	v.e.Add(err)
}

func (v *Validator) Presence(attr string, value string) *PresenceError {
	verr := Presence(attr, value)
	if verr != nil {
		v.Add(verr)
	}

	return verr
}

func (v *Validator) Length(attr string, value string, min, max int) *LengthError {
	verr := Length(attr, value, min, max)
	if verr != nil {
		v.Add(verr)
	}

	return verr
}

func (v *Validator) IsValid() bool {
	return v.e.Len() == 0
}

// Errors returns the validation errors. It will always be nil or a value of type Errors. It purposely returns type
// error instead of Errors to avoid the issue where an interface has a type but a nil value. This can can test to !=
// nil.
func (v *Validator) Errors() error {
	if v.IsValid() {
		return nil
	}

	return v.e
}

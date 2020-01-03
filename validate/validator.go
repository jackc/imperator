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

func (v *Validator) Errors() Errors {
	return v.e
}

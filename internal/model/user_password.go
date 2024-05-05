package model

type Password string

func (p Password) Validate(rules ...func(Password) error) error {
	for _, rule := range rules {
		if err := rule(p); err != nil {
			return err
		}
	}

	return nil
}

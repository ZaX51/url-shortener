package validator

import (
	"fmt"
	"net/url"
	"regexp"
)

type Validator struct {
	rgexp *regexp.Regexp
}

const regexpStr = `^(.*)localhost$`

func New() *Validator {
	validator := new(Validator)
	validator.rgexp = regexp.MustCompile(regexpStr)

	return validator
}

func (v *Validator) ValidateUrl(str string) error {
	if len(str) == 0 {
		return fmt.Errorf("empty URL")
	}

	parsedUrl, err := url.ParseRequestURI(str)
	if err != nil {
		return fmt.Errorf("bad URL")
	}

	res := v.rgexp.MatchString(parsedUrl.Hostname())
	if res {
		return fmt.Errorf("used restricted domain name")
	}

	return nil
}

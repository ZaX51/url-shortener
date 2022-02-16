package validator

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

type Validator struct {
	rgexp *regexp.Regexp
}

func New(domain string) *Validator {
	validator := new(Validator)
	validator.rgexp = regexp.MustCompile(fmt.Sprintf("^(.*)%s$", domain))

	return validator
}

func (v *Validator) ValidateUrl(str string) error {
	if len(str) == 0 {
		return errors.New("empty URL")
	}

	parsedUrl, err := url.ParseRequestURI(str)
	if err != nil {
		return errors.New("bad URL")
	}

	switch parsedUrl.Scheme {
	case "http", "https":
	default:
		return errors.New("invalid scheme")
	}

	res := v.rgexp.MatchString(parsedUrl.Hostname())
	if res {
		return errors.New("used restricted domain name")
	}

	return nil
}

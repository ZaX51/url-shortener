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
		return errors.New("EMPTY_URL")
	}

	parsedUrl, err := url.ParseRequestURI(str)
	if err != nil {
		return errors.New("BAD_URL")
	}

	switch parsedUrl.Scheme {
	case "http", "https":
	default:
		return errors.New("INVALID_SCHEME")
	}

	res := v.rgexp.MatchString(parsedUrl.Host)
	if res {
		return errors.New("RESTRICTED_DOMAIN")
	}

	return nil
}

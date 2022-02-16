package validator_test

import (
	"fmt"
	"testing"

	"github.com/ZaX51/url-shortener/internal/validator"
)

const test_domain = "domain.xyz"

func TestValidatorEmptyInput(t *testing.T) {
	validator := validator.New(test_domain)
	input := ""

	err := validator.ValidateUrl(input)

	if err == nil || err.Error() != "empty URL" {
		t.Errorf("Validator.ValidateUrl should return error on input: '%s'", input)
	}
}

func TestValidatorValidUrl(t *testing.T) {
	validator := validator.New(test_domain)
	input := "http://test.xyz"

	err := validator.ValidateUrl(input)

	if err != nil {
		t.Errorf("Validator.ValidateUrl returned error on input: '%s'", input)
	}
}

func TestValidatorValidHttpsUrl(t *testing.T) {
	validator := validator.New(test_domain)
	input := "https://test.xyz"

	err := validator.ValidateUrl(input)

	fmt.Println(err)

	if err != nil {
		t.Errorf("Validator.ValidateUrl returned error on input: '%s'", input)
	}
}

func TestValidatorInvalidUrl(t *testing.T) {
	validator := validator.New(test_domain)
	input := "test"

	err := validator.ValidateUrl(input)

	fmt.Println(err)

	if err == nil || err.Error() != "bad URL" {
		t.Errorf("Validator.ValidateUrl should return error on input with invalid scheme: '%s'", input)
	}
}

func TestValidatorInvalidUrlScheme(t *testing.T) {
	validator := validator.New(test_domain)
	input := "ht://test.xyz"

	err := validator.ValidateUrl(input)

	fmt.Println(err)

	if err == nil || err.Error() != "invalid scheme" {
		t.Errorf("Validator.ValidateUrl should return error on input with invalid scheme: '%s'", input)
	}
}

func TestValidatorRestrictedDomain(t *testing.T) {
	validator := validator.New(test_domain)
	input := "http://domain.xyz/122"

	err := validator.ValidateUrl(input)

	if err == nil || err.Error() != "used restricted domain name" {
		t.Errorf("Validator.ValidateUrl should return error on input with restricted domain: '%s'", input)
	}
}

func TestValidatorRestrictedSubDomain(t *testing.T) {
	validator := validator.New(test_domain)
	input := "http://subdomain.domain.xyz"

	err := validator.ValidateUrl(input)

	if err == nil || err.Error() != "used restricted domain name" {
		t.Errorf("Validator.ValidateUrl should return error on input with restricted domain: '%s'", input)
	}
}

package validator_test

import (
	"testing"

	"github.com/ZaX51/url-shortener/internal/validator"
	"github.com/stretchr/testify/assert"
)

const test_domain = "domain.xyz"

func TestValidatorEmptyInput(t *testing.T) {
	validator := validator.New(test_domain)
	input := ""

	err := validator.ValidateUrl(input)

	assert.EqualError(t, err, "empty URL")
}

func TestValidatorValidUrl(t *testing.T) {
	validator := validator.New(test_domain)
	input := "http://test.xyz"

	err := validator.ValidateUrl(input)

	assert.NoError(t, err)
}

func TestValidatorValidHttpsUrl(t *testing.T) {
	validator := validator.New(test_domain)
	input := "https://test.xyz"

	err := validator.ValidateUrl(input)

	assert.NoError(t, err)
}

func TestValidatorInvalidUrl(t *testing.T) {
	validator := validator.New(test_domain)
	input := "test"

	err := validator.ValidateUrl(input)

	assert.EqualError(t, err, "bad URL")
}

func TestValidatorInvalidUrlScheme(t *testing.T) {
	validator := validator.New(test_domain)
	input := "ht://test.xyz"

	err := validator.ValidateUrl(input)

	assert.EqualError(t, err, "invalid scheme")
}

func TestValidatorRestrictedDomain(t *testing.T) {
	validator := validator.New(test_domain)
	input := "http://domain.xyz/122"

	err := validator.ValidateUrl(input)

	assert.EqualError(t, err, "used restricted domain name")
}

func TestValidatorRestrictedSubDomain(t *testing.T) {
	validator := validator.New(test_domain)
	input := "http://subdomain.domain.xyz"

	err := validator.ValidateUrl(input)

	assert.EqualError(t, err, "used restricted domain name")
}

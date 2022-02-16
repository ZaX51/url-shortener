package base62_test

import (
	"testing"

	"github.com/ZaX51/url-shortener/internal/base62"
	"github.com/stretchr/testify/assert"
)

func TestToBase64EmptyArray(t *testing.T) {
	result := base62.ToBase62([]byte{})

	if len(result) != 0 {
		t.Errorf("ToBase62 = %s; want empty string", result)
	}
}

func TestToBase64OneItemArray(t *testing.T) {
	result := base62.ToBase62([]byte{0x61})

	expected := "1Z"

	assert.Equal(t, expected, result)
}

func TestToBase64String(t *testing.T) {
	result := base62.ToBase62([]byte("qwertyzaerty"))

	expected := "jfXgVTpHPdtEqTpB"

	assert.Equal(t, expected, result)
}

func TestToBase64UrlString(t *testing.T) {
	result := base62.ToBase62([]byte("https://test.test"))

	expected := "D7qW2GM8nm2wH4FoGdNnDXw"

	assert.Equal(t, expected, result)
}

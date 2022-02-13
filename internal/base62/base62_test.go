package base62_test

import (
	"testing"

	"github.com/ZaX51/url-shortener/internal/base62"
)

func TestTobase64EmptyArray(t *testing.T) {
	result := base62.ToBase62([]byte{})

	if len(result) != 0 {
		t.Errorf("ToBase62 = %s; want empty string", result)
	}
}

func TestTobase64OneItemArray(t *testing.T) {
	result := base62.ToBase62([]byte{0x61})

	expected := "1Z"

	if result != expected {
		t.Errorf("ToBase62 = %s; want '%s' string", result, expected)
	}
}

func TestTobase64String(t *testing.T) {
	result := base62.ToBase62([]byte("qwertyzaerty"))

	expected := "jfXgVTpHPdtEqTpB"

	if result != expected {
		t.Errorf("ToBase62 = %s; want '%s' string", result, expected)
	}
}

func TestTobase64UrlString(t *testing.T) {
	result := base62.ToBase62([]byte("https://test.test"))

	expected := "D7qW2GM8nm2wH4FoGdNnDXw"

	if result != expected {
		t.Errorf("ToBase62 = %s; want '%s' string", result, expected)
	}
}

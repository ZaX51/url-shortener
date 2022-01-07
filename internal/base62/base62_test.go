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

	if result != "1Z" {
		t.Errorf("ToBase62 = %s; want '1Z' string", result)
	}
}

func TestTobase64String(t *testing.T) {
	result := base62.ToBase62([]byte("qwertyzaerty"))

	if result != "jfXgVTpHPdtEqTpB" {
		t.Errorf("ToBase62 = %s; want 'jfXgVTpHPdtEqTpB' string", result)
	}
}

func TestTobase64UrlString(t *testing.T) {
	result := base62.ToBase62([]byte("https://test.test"))

	if result != "D7qW2GM8nm2wH4FoGdNnDXw" {
		t.Errorf("ToBase62 = %s; want 'D7qW2GM8nm2wH4FoGdNnDXw' string", result)
	}
}

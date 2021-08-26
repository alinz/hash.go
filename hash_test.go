package hash_test

import (
	"testing"

	"github.com/alinz/hash.go"
)

func TestValueStringConverstion(t *testing.T) {
	expectedHashValue := "sha256-2498ad992b02c2f6e21684e8057a01463acad5c75a4e75d095619c556a559e8c"
	hashValue, err := hash.ValueFromString(expectedHashValue)
	if err != nil {
		t.Fatal(err)
	}

	if expectedHashValue != hashValue.String() {
		t.Fatalf("expected %s but got this %s", expectedHashValue, hashValue.String())
	}
}

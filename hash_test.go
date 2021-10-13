package hash_test

import (
	"encoding/json"
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

func TestJsonMarshalUnMarshal(t *testing.T) {
	expectedHashValue := "sha256-2498ad992b02c2f6e21684e8057a01463acad5c75a4e75d095619c556a559e8c"
	hashValue, err := hash.ValueFromString(expectedHashValue)
	if err != nil {
		t.Fatal(err)
	}

	jsonHashValue, err := hashValue.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	var hashValueFromJson hash.Value
	err = hashValueFromJson.UnmarshalJSON(jsonHashValue)
	if err != nil {
		t.Fatal(err)
	}

	if expectedHashValue != hashValueFromJson.String() {
		t.Fatalf("expected %s but got this %s", expectedHashValue, hashValueFromJson.String())
	}
}

func TestStructMarshalUnmarshal(t *testing.T) {
	type testStruct struct {
		HashValue hash.Value `json:"hash_value"`
	}

	expectedHashValue := hash.Bytes([]byte("test"))

	b, err := json.Marshal(testStruct{
		HashValue: expectedHashValue,
	})
	if err != nil {
		t.Fatal(err)
	}

	var result testStruct
	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.HashValue.String() != expectedHashValue.String() {
		t.Fatalf("expected %s but got this %s", expectedHashValue, result.HashValue.String())
	}
}

package hash

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"strings"
)

const (
	hashName   = "sha256"
	hashHeader = hashName + "-"
)

// Value is a custom type which provides formating hash values with ease
type Value []byte

// String returns the full string version of hash along with header
func (v *Value) String() string {
	return fmt.Sprintf("%s-%x", hashName, []byte(*v))
}

func (v *Value) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, `"`)
	value, err := ValueFromString(string(b))
	if err != nil {
		return err
	}

	*v = value

	return nil
}

func (v *Value) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", v.String())), nil
}

// Short returns the last 5 characters of hash
func (v *Value) Short() string {
	val := v.String()
	return val[len(val)-5:]
}

// Bytes hashes the given content in bytes and
// produces sha256
func Bytes(content []byte) Value {
	hasher := sha256.New()
	hasher.Write(content)
	return hasher.Sum(nil)
}

type Reader struct {
	r      io.Reader
	hasher hash.Hash
}

func (r *Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if n > 0 {
		r.hasher.Write(b[:n])
	} else if n == 0 {
		return 0, io.EOF
	}

	return n, err
}

// Hash generates the hash value of the read content
// NOTE: make sure to call this function once you finish reading it
func (r Reader) Hash() Value {
	return r.hasher.Sum(nil)
}

// NewReader creates parallel hashing with reading
// it is usuful if you want to read the reader and at the same time
// calculate the hash value of the read content
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r:      r,
		hasher: sha256.New(),
	}
}

// ValueFromString tries to parse the string version of value
// which was produced by calling String() methid back to type Value
func ValueFromString(value string) (Value, error) {
	return hex.DecodeString(strings.Replace(value, hashHeader, "", 1))
}

// Print a handy function similar to fmt.Fprint which accept
// hash value as second arguments
func Print(w io.Writer, hash []byte, args ...interface{}) {
	value := Value(hash)
	args = append([]interface{}{value.Short()}, args...)
	fmt.Fprintln(w, args...)
}

// Format a handy function to convert value in bytes in string
func Format(value []byte) string {
	if value == nil {
		return "nil"
	}

	v := Value(value)
	return v.String()
}

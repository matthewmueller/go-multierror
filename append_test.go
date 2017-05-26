package multierror

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestAppend_Error(t *testing.T) {
	original := &Error{
		Errors: []error{errors.New("foo")},
	}

	result := Append(original, errors.New("bar"))
	if len(result.Errors) != 2 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}

	original = &Error{}
	result = Append(original, errors.New("bar"))
	if len(result.Errors) != 1 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}

	// Test when a typed nil is passed
	var e *Error
	result = Append(e, errors.New("baz"))
	if len(result.Errors) != 1 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}

	// Test flattening
	original = &Error{
		Errors: []error{errors.New("foo")},
	}

	result = Append(original, Append(nil, errors.New("foo"), errors.New("bar")))
	if len(result.Errors) != 3 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}
}

func TestAppend_NilError(t *testing.T) {
	var err error
	result := Append(err, errors.New("bar"))
	if len(result.Errors) != 1 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}
}

func TestAppend_NilErrorArg(t *testing.T) {
	var err error
	var nilErr *Error
	result := Append(err, nilErr)
	if result != nil {
		t.Fatalf("result is not nil: %s", result.Error())
	}
}

func TestAppend_NilErrorIfaceArg(t *testing.T) {
	var err error
	var nilErr error
	result := Append(err, nilErr)
	if result != nil {
		t.Fatalf("result is not nil: %s", result.Error())
	}
}

func TestAppend_Nils(t *testing.T) {
	var err error
	var err1 error
	// var err2 error
	err = Append(err, err1)
	if result != nil {
		fmt.Println("WTF", reflect.TypeOf(err))
	}
	// if err != nil {
	// 	t.Fatalf("1. result is not nil: %s", err)
	// }
	// err = Append(err, err2)
	// if err != nil {
	// 	t.Fatalf("2. result is not nil: %s", err)
	// }
}

func TestAppend_NonError(t *testing.T) {
	original := errors.New("foo")
	result := Append(original, errors.New("bar"))
	if len(result.Errors) != 2 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}
}

func TestAppend_NonError_Error(t *testing.T) {
	original := errors.New("foo")
	result := Append(original, Append(nil, errors.New("bar")))
	if len(result.Errors) != 2 {
		t.Fatalf("wrong len: %d", len(result.Errors))
	}
}

package tests

import (
	"cmp_lab/src/structs/opt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomeOption(t *testing.T) {
	value := 42
	option := opt.Some(value)

	if !option.IsSome() {
		t.Error("Option should be Some")
	}

	if option.IsNone() {
		t.Error("Option should not be None")
	}

	if option.Unwrap() != value {
		t.Errorf("Expected value %v, but got %v", value, option.Unwrap())
	}
}

func TestNoneOption(t *testing.T) {
	option := opt.None[int]()

	if option.IsSome() {
		t.Error("Option should not be Some")
	}

	if !option.IsNone() {
		t.Error("Option should be None")
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Expected panic, but none occurred")
		}
	}()

	option.Unwrap()
}

func TestOption(t *testing.T) {
	// Test Some function
	var value int = 42
	some := opt.Some(value)
	assert.NotNil(t, some.Value, "Some option should not be nil")

	// Test None function
	none := opt.None[int]()
	assert.Nil(t, none.Value, "None option should be nil")

	// Test IsSome and IsNone function
	assert.True(t, some.IsSome(), "Option should be Some")
	assert.False(t, some.IsNone(), "Option should not be None")
	assert.True(t, none.IsNone(), "Option should be None")
	assert.False(t, none.IsSome(), "Option should not be Some")

	// Test Unwrap function
	unwrapped := some.Unwrap()
	assert.Equal(t, value, unwrapped, "Unwrapped value should match input")

	// Test ValueOrCall function
	var called bool
	callback := func() { called = true }
	valueOrCall := some.ValueOrCall(callback)
	assert.Equal(t, value, valueOrCall, "ValueOrCall should return value")
	assert.False(t, called, "Callback should not be called when option is Some")

	called = false
	valueOrCall = none.ValueOrCall(callback)
	assert.Equal(t, 0, valueOrCall, "ValueOrCall should return zero value for type")
	assert.True(t, called, "Callback should be called when option is None")

	// Test ValueOr function
	valueOr := some.ValueOr(0)
	assert.Equal(t, value, valueOr, "ValueOr should return value")
	valueOr = none.ValueOr(0)
	assert.Equal(t, 0, valueOr, "ValueOr should return default value when option is None")

	// Test panic when calling Unwrap on None option
	assert.Panics(t, func() { none.Unwrap() }, "Unwrap should panic when option is None")
}

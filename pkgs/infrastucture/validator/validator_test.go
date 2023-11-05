package validator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Field string `binding:"required"`
}

func TestValidateStruct(t *testing.T) {
	v := &DefaultValidator{}

	err := v.ValidateStruct(testStruct{})
	assert.NotNil(t, err)

	err = v.ValidateStruct(&testStruct{Field: "value"})
	assert.Nil(t, err)

	err = v.ValidateStruct("not a struct")
	assert.Nil(t, err)
}

func TestEngine(t *testing.T) {
	v := &DefaultValidator{}
	engine := v.Engine()
	_, ok := engine.(*validator.Validate)
	assert.True(t, ok)
}

func TestValidate(t *testing.T) {
	errors := Validate(testStruct{})
	assert.NotEmpty(t, errors)

	errors = Validate(&testStruct{Field: "value"})
	assert.Empty(t, errors)
}

func TestVal(t *testing.T) {
	err := val(testStruct{})
	assert.NotNil(t, err)

	err = val(&testStruct{Field: "value"})
	assert.Nil(t, err)
}

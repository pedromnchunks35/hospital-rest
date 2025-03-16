package validation

import (
	"github.com/stretchr/testify/assert"
	"hospital-rest/utils/validation"
	"testing"
)

const pageString = "0"
const sizeString = "1"
const pageExpected = 0
const sizeExpected = 1
const badNumberPageString = "-1"
const badPageStringChar = "l"
const badNumberSizeString = "-1"
const badSizeStringChar = "l"
const sizeZeroString = "0"

func Test_GivenValidPageStringAndSize_whenValidate_thenReturnValidObject(t *testing.T) {
	page, size, err := validation.ValidatePageAndSize(pageString, sizeString)

	assert.Nil(t, err)
	assert.Equal(t, page, pageExpected)
	assert.Equal(t, size, sizeExpected)
}

func Test_GivenBadNumberPageStringAndValidSize_whenValidate_thenReturnError(t *testing.T) {
	_, _, err := validation.ValidatePageAndSize(badNumberPageString, sizeString)

	assert.NotNil(t, err)
}

func Test_GivenValidPageAndInvalidSizeString_whenValidate_thenReturnError(t *testing.T) {
	_, _, err := validation.ValidatePageAndSize(pageString, badNumberSizeString)

	assert.NotNil(t, err)
}

func Test_GivenCharInPageAndValidSizeString_whenValidate_thenReturnError(t *testing.T) {
	_, _, err := validation.ValidatePageAndSize(badPageStringChar, sizeString)

	assert.NotNil(t, err)
}

func Test_GivenValidPageAndCharInSize_whenValidate_thenReturnError(t *testing.T) {
	_, _, err := validation.ValidatePageAndSize(pageString, badSizeStringChar)

	assert.NotNil(t, err)
}

func Test_GivenValidPageButSizeZero_whenValidate_thenReturnError(t *testing.T) {
	_, _, err := validation.ValidatePageAndSize(pageString, sizeZeroString)

	assert.NotNil(t, err)
}

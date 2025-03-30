package validation

import (
	"fmt"
	"hospital-rest/dtos"
	"strconv"
)

func ValidatePageAndSize(pageString string, sizeString string) (int, int, error) {
	page, err := convertStringToInt(pageString)
	if err != nil {
		return 0, 0, err
	}

	size, err := convertStringToInt(sizeString)
	if err != nil {
		return 0, 0, err
	}

	bothLegit := arePageAndSizeLegit(page, size)
	if !bothLegit {
		return 0, 0, fmt.Errorf("page and size are not consistent")
	}

	return page, size, nil
}

func ValidateObject(asset *dtos.PostAssetRequest) error {
	if isStringValid(asset.Id) &&
		isStringValid(asset.TypeForm) &&
		isStringValid(asset.Hash) &&
		isStringValid(asset.Description) &&
		isStringValid(asset.InsertionType) {
		return nil
	}
	return fmt.Errorf("please make sure that all the fields are present")
}

func isStringValid(obj string) bool {
	return len(obj) != 0
}

func arePageAndSizeLegit(page int, size int) bool {
	return !isNumberNegative(page) && !isNumberNegative(size) && isNumberDifferentThatZero(size)
}

func convertStringToInt(toConvert string) (int, error) {
	return strconv.Atoi(toConvert)
}

func isNumberNegative(number int) bool {
	return number < 0
}

func isNumberDifferentThatZero(number int) bool {
	return number != 0
}

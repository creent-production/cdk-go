package parser

import (
	"regexp"
	"strconv"
)

func ParsePathToInt(regex, path string) (int, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return 0, err
	}
	values := re.FindStringSubmatch(path)
	result, err := strconv.Atoi(values[1])
	if err != nil {
		return 0, err
	}
	return result, nil
}

func ParseSliceIntToSliceStr(values []int) []string {
	valuesText := []string{}

	for i := range values {
		text := strconv.Itoa(values[i])
		valuesText = append(valuesText, text)
	}

	return valuesText
}

func ParseSliceStrToSliceInt(values []string) ([]int, error) {
	valuesInt := []int{}

	for i := range values {
		integer, err := strconv.Atoi(values[i])
		if err != nil {
			return []int{}, err
		}
		valuesInt = append(valuesInt, integer)
	}

	return valuesInt, nil
}

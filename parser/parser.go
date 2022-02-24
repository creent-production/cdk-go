package parser

import (
	"regexp"
	"strconv"
	"strings"
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

func ParseSliceUint8ToSliceInt(values []uint8) ([]int, error) {
	r := regexp.MustCompile(`[0-9,]`)
	matches := r.FindAllString(string(values), -1)

	valuesInt := []int{}
	data := strings.Split(strings.Join(matches, ""), ",")
	for _, i := range data {
		integer, err := strconv.Atoi(i)
		if err != nil {
			return []int{}, err
		}
		valuesInt = append(valuesInt, integer)
	}

	return valuesInt, nil
}

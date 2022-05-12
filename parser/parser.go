package parser

import (
	"regexp"
	"sort"
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

func ParsePathToStr(regex, path string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	values := re.FindStringSubmatch(path)

	return values[1], nil
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

func ParseSliceUint8ToSliceStr(values []uint8) []string {
	r := regexp.MustCompile(`[^{}]+`)
	matches := r.FindAllString(string(values), -1)

	return strings.Split(strings.Join(matches, ""), ",")
}

func RemoveIndexStr(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func CompareTwoSliceStr(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	// sort slice
	sort.Strings(one)
	sort.Strings(two)

	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}

func CompareTwoSliceInt(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}
	// sort slice
	sort.Ints(one)
	sort.Ints(two)

	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}

func DifferenceStrFoundOrNot(a, b []string, f bool) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if f {
			if _, found := mb[x]; found {
				diff = append(diff, x)
			}
		} else {
			if _, found := mb[x]; !found {
				diff = append(diff, x)
			}
		}
	}
	return diff
}

func DifferenceIntFoundOrNot(a, b []int, f bool) []int {
	mb := make(map[int]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int
	for _, x := range a {
		if f {
			if _, found := mb[x]; found {
				diff = append(diff, x)
			}
		} else {
			if _, found := mb[x]; !found {
				diff = append(diff, x)
			}
		}
	}
	return diff
}

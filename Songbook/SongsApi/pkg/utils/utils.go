package utils

import "strconv"

func CheckOffset(offset string) (int, error) {
	result, err := strconv.Atoi(offset)
	if err != nil {
		return -1, err
	} else {
		return result, nil
	}
}

func CheckLimit(limit string) (int, error) {
	result, err := strconv.Atoi(limit)
	if err != nil {
		return -1, err
	} else {
		return result, nil
	}
}

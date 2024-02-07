package utils

import "strconv"

func ConvertStringToIntSlice(values ...string) (ret []int, err error) {
	for _, value := range values {
		var intValue int

		if intValue, err = strconv.Atoi(value); err != nil {
			return
		}

		ret = append(ret, intValue)
	}

	return
}

func EmptyArray(length int) []byte {
	var array []byte
	for i := 0; i < length; i++ {
		array = append(array, []byte(" ")...)
	}
	return array
}

func PtrAny[T any](value T) *T {
	return &value
}

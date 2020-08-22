package service

import (
	"fmt"
	"reflect"
)

// AllowValue checks if the allowed values are selected.
func AllowValue(chars []string) error {
	// Loop through charFlags to check if it contains any unacceptable characters.
	for _, ch := range chars {
		bool, err := Contains(ch, []string{"l", "u", "n", "s", "c"})
		if err != nil {
			return err
		}
		if !bool {
			return fmt.Errorf("invalid \"%v\" for flag -char", ch)
		}
	}

	return nil
}

// Contains checks if a value exists in slice.
func Contains(target interface{}, list interface{}) (bool, error) {
	switch list.(type) {
	default:
		return false, fmt.Errorf("%v is an unsupported type", reflect.TypeOf(list))
	case []int:
		revert := list.([]int)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil

	case []uint64:
		revert := list.([]uint64)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil

	case []string:
		revert := list.([]string)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil
	}
}

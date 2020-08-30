package slice

import "fmt"

// スライスの中に指定した要素が含まれているかを判定する
func Contains(target interface{}, list interface{}) (bool, error) {
	switch list.(type) {
	case []string:
		l := list.([]string)
		for _, v := range l {
			if v == target {
				return true, nil
			}
		}
	default:
		return false, fmt.Errorf("unsupported format")
	}
	return false, nil
}

package util

import "strconv"

func NumericToString(id interface{}) string {
	switch v := id.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case string:
		return v
	default:
		return ""
	}
}

func ReverseData(data []interface{}) []interface{} {
	l := len(data)

	reversed := make([]interface{}, l)

	for i := 0; i < len(data); i++ {
		reversed[i] = data[l - 1]

		l--
	}

	return reversed
}
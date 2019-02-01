package exporters

import "strconv"

func boolToString(val bool) string {
	if val {
		return "1"
	}
	return "0"
}

func intToString(val int) string {
	return strconv.FormatInt(int64(val), 10)
}
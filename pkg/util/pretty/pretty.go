package pretty

import "encoding/json"

// Only for normal logging purpose, 4 space indent
func PrettifyJson(i interface{}, indent bool) string {
	var str []byte
	if indent {
		str, _ = json.MarshalIndent(i, "", "    ")
	} else {
		str, _ = json.Marshal(i)
	}
	return string(str)
}
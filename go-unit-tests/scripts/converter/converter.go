package converter

import (
	"encoding/json"
)

func StructToMap(data any) map[string]any {
	b, _ := json.Marshal(&data)

	var m map[string]any
	_ = json.Unmarshal(b, &m)

	return m
}
package utils

import "encoding/json"

// MapToStruct converte um map[string]interface{} em uma struct qualquer.
func MapToStruct(data map[string]interface{}, out interface{}) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}

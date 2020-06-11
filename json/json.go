package json

import (
	JSONIter "github.com/json-iterator/go"
)

var (
	json          = JSONIter.ConfigCompatibleWithStandardLibrary
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)

func ToJSON(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}
	if w, ok := obj.(string); ok {
		var js map[string]interface{}
		if err := json.Unmarshal([]byte(w), &js); err != nil {
			return w
		}
		return js
	}
	res, _ := json.Marshal(obj)
	return string(res)
}

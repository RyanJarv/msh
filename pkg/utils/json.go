package utils

import "strings"

func FixDynamoJsonForGo(event map[string]any) map[string]any {
	result := map[string]any{}
	for key, old := range event {
		updated := map[string]any{}

		for typeName, value := range old.(map[string]any) {
			typeName = strings.ToUpper(typeName)
			if typeName == "NUL" {
				typeName = "NULL"
			}

			if v, ok := value.(map[string]any); ok {
				updated[typeName] = FixDynamoJsonForGo(v)
			} else {
				updated[typeName] = value
			}
		}

		result[key] = updated

	}

	return result
}

func FixDynamoJsonForJava(event map[string]any) map[string]any {
	result := map[string]any{}
	for key, old := range event {
		updated := map[string]any{}

		for typeName, value := range old.(map[string]any) {
			typeName = strings.ToUpper(typeName[0:1]) + strings.ToLower(typeName[1:])
			if typeName == "Nul" {
				typeName = "Null"
			}

			if v, ok := value.(map[string]any); ok {
				updated[typeName] = FixDynamoJsonForGo(v)
			} else {
				updated[typeName] = value
			}
		}

		result[key] = updated

	}

	return result
}

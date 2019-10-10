package utils

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func ToStringPointer(v interface{}) *string {
	r := v.(string)
	return &r
}

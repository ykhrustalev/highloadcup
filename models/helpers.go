package models

func GetNonNullString(value interface{}) (string, error) {
	if value == nil {
		return "", ErrorNullValue
	}

	strValue, ok := value.(string)
	if !ok {
		return "", ErrorNotAString
	}

	return strValue, nil
}

func GetNonNullStringP(value interface{}) (*string, error) {
	val, err := GetNonNullString(value)
	return &val, err
}

func GetNonNullInt(value interface{}) (int, error) {
	if value == nil {
		return 0, ErrorNullValue
	}

	float64Value, ok := value.(float64)
	if ok {
		return int(float64Value), nil
	}

	intValue, ok := value.(int)
	if ok {
		return intValue, nil
	}

	int32Value, ok := value.(int32)
	if ok {
		return int(int32Value), nil
	}

	int64Value, ok := value.(int64)
	if ok {
		return int(int64Value), nil
	}

	return 0, ErrorNotAnInt
}

func GetNonNullIntP(value interface{}) (*int, error) {
	val, err := GetNonNullInt(value)
	return &val, err
}

func GetNonNullInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, ErrorNullValue
	}

	float64Value, ok := value.(float64)
	if ok {
		return int64(float64Value), nil
	}

	intValue, ok := value.(int)
	if ok {
		return int64(intValue), nil
	}

	int32Value, ok := value.(int32)
	if ok {
		return int64(int32Value), nil
	}

	int64Value, ok := value.(int64)
	if ok {
		return int64Value, nil
	}

	return 0, ErrorNotAnInt
}

func GetNonNullInt64P(value interface{}) (*int64, error) {
	val, err := GetNonNullInt64(value)
	return &val, err
}

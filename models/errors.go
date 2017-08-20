package models

import "errors"

var ErrorNullValue = errors.New("field contains null")
var ErrorNotAString = errors.New("field is not a string")
var ErrorNotAnInt = errors.New("field is not int")
var ErrorInvalidGender = errors.New("gender unexpected")

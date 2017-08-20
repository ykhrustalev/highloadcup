package models

import "errors"

var ErrorNullValue = errors.New("field contains null")
var ErrorNotAString = errors.New("field is not a string")
var ErrorNotAnInt = errors.New("field is not int")
var ErrorInvalidGender = errors.New("gender unexpected")
var ErrorStringTooLong = errors.New("string is too long")
var ErrorBirthDayToEarly = errors.New("birthday too early")
var ErrorBirthDayToLate = errors.New("birthday too late")

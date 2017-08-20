package models

import "errors"

var ErrorInvalidId = errors.New("invalid id")
var ErrorNullValue = errors.New("field contains null")
var ErrorNotAString = errors.New("field is not a string")
var ErrorNotAnInt = errors.New("field is not int")
var ErrorInvalidGender = errors.New("gender unexpected")
var ErrorStringOutOfRange = errors.New("string size out of range")
var ErrorBirthDayToEarly = errors.New("birthday too early")
var ErrorBirthDayToLate = errors.New("birthday too late")
var ErrorMark = errors.New("invalid mark")

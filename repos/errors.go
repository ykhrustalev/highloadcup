package repos

import "errors"

var ErrorNotFound = errors.New("missing item")
var ErrorMalformed = errors.New("malformed data")
var ErrorObjectExists = errors.New("object with id exists")

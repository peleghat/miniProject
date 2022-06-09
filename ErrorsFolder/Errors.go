package ErrorsFolder

import (
	"errors"
	_ "errors"
	_ "fmt"
)

var ErrDbConnection = errors.New("failed to connect to the db")
var ErrAlreadyExist = errors.New("the value already exist in the table")
var ErrIllegalValues = errors.New("contains illegal values")

func errorHandler(e error) {

}
package division

/*
# ============================
# Author: Nitika Mishra
# ============================
*/

import "errors"

const (
	DIVISION_BY_ZERO                 = "denominator can't be set to zero"
	INPUT_TYPE_INT64                 = "integer expected value input type int64"
	INPUT_TYPE_INT32                 = "integer expected value input type int32"
	INPUT_TYPE_FLOAT32               = "integer expected value input type float32"
	INPUT_TYPE_FLOAT64               = "integer expected value input type float64"
	INPUT_TYPE_BOOL                  = "integer expected value input type bool"
	INPUT_TYPE_STRING                = "integer expected value input type string"
	INTERFACE_CAST_ERROR			 = "cannot convert interface{} to interger"
)

func getDivisionByZeroError() error {
    return errors.New(DIVISION_BY_ZERO)
}

func getInputInt64Error() error {
    return errors.New(INPUT_TYPE_INT64)
}

func getInputInt32Error() error {
    return errors.New(INPUT_TYPE_INT32)
}

func getInputFloat32Error() error {
    return errors.New(INPUT_TYPE_FLOAT32)
}

func getInputFloat64Error() error {
    return errors.New(INPUT_TYPE_FLOAT64)
}

func getInputStringError() error {
    return errors.New(INPUT_TYPE_STRING)
}

func getInterfaceCastError() error {
	return errors.New(INTERFACE_CAST_ERROR)
}

func getInputBoolError() error {
	return errors.New(INPUT_TYPE_BOOL)
}
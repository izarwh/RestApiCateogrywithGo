package helper

import "prrestapi/exception"

func PanicIfError(err error) {
	if err != nil {
		panic(exception.NewInternalServerErr(err))
	}
}

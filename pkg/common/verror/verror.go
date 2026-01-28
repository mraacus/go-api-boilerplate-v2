package verror

import "errors"

var (
	ServiceInternalError  = errors.New("ServiceInternalError")
	RequestParamsError    = errors.New("RequestParamsError")
	ParamsAbsent          = errors.New("ParamsAbsent")
	DataDoesNotExist      = errors.New("DataDoesNotExist")
)
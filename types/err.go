package spec

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	ORIGIN           Code = 1000
	JSON_MARSHAL_ERR Code = iota + 1000
	JSON_UNMARSHAL_ERR
	PARSEMULTIPART_ERR
	IOUTIL_READALL_ERR
	IO_COPYN_ERR
	SCHEDULE_ERR
	CID_PARSE_ERR
	KUBERBETES_API_FAIL
	SECTOR_TYPE_ERR
	OS_ERR
	WRITE_ERR
	PARSEUINT_ERR
	PARSEINT_ERR
	PROVIDE_RESP_ERR
)

const FILTAB_SPLITER = "<filtab>"

type FiltabErr struct {
	Err  error
	Code Code
}

type Code int64

func (c Code) String() string {
	return strconv.FormatInt(int64(c), 10)
}

/**
ErrorCode映射关系
*/
func CodeStr(code Code) string {
	switch code {
	case ORIGIN:
		return "origin sealer ffi error"
	case JSON_MARSHAL_ERR:
		return "json marshal error"
	case JSON_UNMARSHAL_ERR:
		return "json unmarshal error"
	case PARSEMULTIPART_ERR:
		return "parsemultipart error"
	case IOUTIL_READALL_ERR:
		return "ioutil readall error"
	case IO_COPYN_ERR:
		return "io copyn error"
	case SCHEDULE_ERR:
		return "schedule error"
	case CID_PARSE_ERR:
		return "cid parse error"
	case KUBERBETES_API_FAIL:
		return "k8s error"
	case SECTOR_TYPE_ERR:
		return "sector type error"
	case OS_ERR:
		return "OS error"
	case WRITE_ERR:
		return "write error"
	case PARSEUINT_ERR:
		return "parse uint error"
	case PARSEINT_ERR:
		return "parse int error"
	case PROVIDE_RESP_ERR:
		return "provide resp error"
	}

	panic("code not found")
	return ""
}

/**
拼接FiltabErr
*/
func NewFiltabErr(code Code, err error) FiltabErr {
	return FiltabErr{
		Err:  err,
		Code: code,
	}
}

/**
从string转换为error
this may introduce panic if string not matched to FiltabErr
*/
func ErrFromString(transErr string) FiltabErr {
	countSplit := strings.Split(transErr, FILTAB_SPLITER)
	if len(countSplit) != 3 {
		panic("error type not FiltabErr")
	}
	errorCode, err := strconv.ParseInt(countSplit[0], 10, 64)
	if err != nil {
		errorCode = 0
	}
	return FiltabErr{
		Err:  errors.New(countSplit[2]),
		Code: Code(errorCode),
	}
}

/**
转换为error string
*/
func (f FiltabErr) ToString() string {
	return f.Code.String() + FILTAB_SPLITER + CodeStr(f.Code) + FILTAB_SPLITER + f.Err.Error()
}

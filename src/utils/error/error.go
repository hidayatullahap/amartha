package error

import (
	"errors"
	"net/http"
)

var (
	ErrGeneral      = errors.New("terdapat kesalahan pada sistem")
	ErrDataNotFound = errors.New("data tidak ditemukan")
	ErrUnauthorized = errors.New("unauthorized")
)

var httpStatusToErrors = map[int][]error{
	http.StatusInternalServerError: {
		ErrGeneral,
	},
	http.StatusNotFound: {
		ErrDataNotFound,
	},
	http.StatusUnauthorized: {
		ErrUnauthorized,
	},
}

func GetHttpCodeByError(err error) int {
	for status, errs := range httpStatusToErrors {
		for _, e := range errs {
			if errors.Is(err, e) {
				return status
			}
		}
	}
	return http.StatusInternalServerError
}

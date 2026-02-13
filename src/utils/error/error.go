package error

import (
	"database/sql"
	"errors"
	"net/http"
)

var (
	ErrGeneral                = errors.New("terdapat kesalahan pada sistem")
	ErrDataNotFound           = errors.New("data tidak ditemukan")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrInvalidLoanStateUpdate = errors.New("tidak bisa melakukan perubahan status")
)

var httpStatusToErrors = map[int][]error{
	http.StatusInternalServerError: {
		ErrGeneral,
	},
	http.StatusNotFound: {
		ErrDataNotFound,
		sql.ErrNoRows,
	},
	http.StatusPreconditionFailed: {
		ErrInvalidLoanStateUpdate,
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

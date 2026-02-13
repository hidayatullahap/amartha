package error

import (
	"database/sql"
	"errors"
	"net/http"

	"modernc.org/sqlite"
)

var (
	ErrGeneral                = errors.New("terdapat kesalahan pada sistem")
	ErrDataNotFound           = errors.New("data tidak ditemukan")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrInvalidLoanStateUpdate = errors.New("tidak bisa melakukan perubahan status")
	ErrExceedsPrincipleAmount = errors.New("investment exceeds loan principal amount")
)

var httpStatusToErrors = map[int][]error{
	http.StatusInternalServerError: {
		ErrGeneral,
	},
	http.StatusNotFound: {
		ErrDataNotFound,
		sql.ErrNoRows,
	},
	http.StatusBadRequest: {
		ErrExceedsPrincipleAmount,
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

	var sqliteErr *sqlite.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code() == 1555 {
			return http.StatusConflict
		}
	}

	return http.StatusInternalServerError
}

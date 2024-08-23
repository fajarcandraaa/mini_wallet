package dto

import (
	"github.com/fajarcandraaa/mini_wallet/internal/presentation"
)

func ToResponse(status string, data interface{}) presentation.Response {
	res := presentation.Response{
		Status: status,
		Data:   data,
	}

	return res
}

package utils

import (
	"math"
	"net/http"
)
type paginationInfo struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
	TotalPages int64 `json:"totalPages"`
	TotalItems int64 `json:"totalItems"`
}

type dataWithPagination  struct {
	Data any `json:"data"`
	Pagination paginationInfo `json:"pagination"`
}

func SendResponseWithPagination(w http.ResponseWriter, data any, page, limit, count int64,) {

	totalPages := math.Ceil(float64(count)/float64(limit))
	totalItems := count

	paginationInfo := paginationInfo{
		Page:page,
		Limit:limit,
		TotalItems: totalItems,
		TotalPages: int64(totalPages),
	}
	dataWithPagination := dataWithPagination{
		Data:data,
		Pagination: paginationInfo,
	}

	SendJSONResponse(w,http.StatusOK,"Successfully fetched", dataWithPagination)
}
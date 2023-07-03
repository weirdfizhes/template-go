package echohttp

import (
	"math"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

type ResponsePaginate struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Error    string      `json:"error"`
	paginate `json:"paginate,omitempty"`
}

type paginate struct {
	CurrentPage int   `json:"current_page,omitempty"`
	Limit       int   `json:"limit,omitempty"`
	TotalPage   int64 `json:"total_page,omitempty"`
	TotalData   int64 `json:"total_data,omitempty"`
}

// ResponseData for handling response data (error and success) without pagination
func ResponseData(ctx echo.Context, code int, message string, data interface{}, err error) error {
	var errorString string

	if err != nil {
		errorString = err.Error()
	} else {
		errorString = ""
	}

	return ctx.JSONPretty(code, Response{
		Message: message,
		Data:    data,
		Error:   errorString,
	}, "")
}

// ResponseData for handling response data (error and success) with pagination
func ResponsePagination(ctx echo.Context, code int, message string, data interface{}, err error, page int, limit int, totalData int64) error {
	var (
		totalPage   int64
		errorString string
	)

	if err != nil {
		errorString = err.Error()
	} else {
		errorString = ""
	}

	if totalData != 0 {
		if page != 0 && limit != 0 {
			total := float64(totalData) / float64(limit)
			if total != 0 {
				totalPage = int64(math.Round(total))
			} else {
				totalPage = 1
			}

			if total > float64(totalPage) {
				totalPage += 1
			}

		} else {
			totalPage = 1
			page = 1
			limit = int(totalData)
		}
	} else {
		totalPage = 1
		page = 1
		limit = 1
	}

	return ctx.JSONPretty(code, ResponsePaginate{
		Message: message,
		Data:    data,
		Error:   errorString,
		paginate: paginate{
			CurrentPage: page,
			Limit:       limit,
			TotalPage:   totalPage,
			TotalData:   totalData,
		},
	}, "")
}

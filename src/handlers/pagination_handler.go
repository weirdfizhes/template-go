package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Pagination struct {
	Limit  int `json:"limit"`
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

func PaginationCountHandler(c echo.Context) (paginate Pagination, err error) {
	limitString := c.QueryParam("limit")
	if limitString == "" {
		limitString = "0"
	}

	pageString := c.QueryParam("page")
	if pageString == "" {
		pageString = "0"
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return
	}

	offset := limit * (page - 1)

	return Pagination{
		Limit:  limit,
		Page:   page,
		Offset: offset,
	}, err
}

package coupon

import (
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
)

func bindQuery(c echo.Context) *FindListCouponDto {
	var query FindListCouponDto
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	if p, err := strconv.Atoi(page); err != nil {
		query.Page = 1
	} else {
		query.Page = p
	}

	if l, err := strconv.Atoi(limit); err != nil {
		query.Limit = 10
	} else {
		query.Limit = l
	}

	query.SortBy = c.QueryParam("sort_by")
	query.OrderBy = c.QueryParam("order_by")

	initializeQuery(&query)

	return &query
}

func initializeQuery(query *FindListCouponDto) {
	sortBys := []string{"discount", "expiry_date", "created_at"}

	if query.OrderBy != "asc" && query.OrderBy != "desc" {
		query.OrderBy = "asc"
	}

	if query.SortBy == "" || slices.Contains(sortBys, query.SortBy) == false {
		query.SortBy = "created_at"
	}
}

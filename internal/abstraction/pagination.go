package abstraction

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

const QUERY_PAGE = "page"
const QUERY_PER_PAGE = "per_page"

type Pagination struct {
	Page     *int    `query:"page" json:"page"`
	PageSize *int    `query:"page_size" json:"page_size"`
	SortBy   *string `query:"sort_by" json:"sort_by"`
	Sort     *string `query:"sort" json:"sort"`
}

func NewFromRequest(c echo.Context) Pagination {
	page, err := strconv.Atoi(c.QueryParam(QUERY_PAGE))
	if err != nil {
		page = 0
	}

	perPage, err := strconv.Atoi(c.QueryParam(QUERY_PER_PAGE))
	if err != nil {
		perPage = 20
	}

	return Pagination{
		Page:     &page,
		PageSize: &perPage,
	}
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
}

func (p *Pagination) Offset() int {
	if *p.Page == 0 {
		return 0 * p.Limit()
	}
	return (*p.Page - 1) * p.Limit()
}

func (p *Pagination) Limit() int {
	return *p.PageSize
}

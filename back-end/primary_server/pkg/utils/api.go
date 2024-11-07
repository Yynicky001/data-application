package utils

import "strconv"

func GetPageQuery(pageParam, perPageParam string) (int, int, error) {
	// 默认第一页
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return 0, 0, err
	}

	// 默认每页10条数据
	if perPageParam == "" {
		perPageParam = "10"
	}
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil {
		return 0, 0, err
	}
	return page, perPage, nil
}

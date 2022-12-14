package utils

import (
	_const "jetdev-task/shared/utils/const"
	"strconv"
)

func PageAttributes(pageStr string) (int, int, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, err
	}

	if err != nil || page <= 0 {
		page = _const.PageNo
	}

	size := _const.PerPageLimit

	return page, size, nil
}

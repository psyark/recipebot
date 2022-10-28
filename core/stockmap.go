package core

import (
	"regexp"
)

type StockMap map[*regexp.Regexp]string

func (sm StockMap) Get(stockName string) (string, bool) {
	for regex, pageID := range sm {
		if regex.MatchString(stockName) {
			return pageID, true
		}
	}
	return "", false
}

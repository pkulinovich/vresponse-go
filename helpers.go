package response

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func buildFirstPath(path string, limit uint64) *string {
	if limit == 0 {
		return &path
	}
	return changePageNumber(path, 1)
}

func buildLastPath(path string, limit, total uint64) *string {
	if limit < 1 {
		return nil
	}
	return changePageNumber(path, getLastPageNumber(limit, total))
}

func buildNextPath(path string, page, limit, total uint64) *string {
	if limit == 0 {
		return &path
	}
	lastPageNumber := getLastPageNumber(limit, total)
	if page >= lastPageNumber {
		return nil
	}
	return changePageNumber(path, page+1)
}

func buildSelfPath(path string, page, limit uint64) *string {
	if limit == 0 {
		return &path
	}
	return changePageNumber(path, page)
}

func buildPrevPath(path string, page, limit, total uint64) *string {
	if page <= 1 {
		return nil
	}
	lastPageNumber := getLastPageNumber(limit, total)
	if page > lastPageNumber {
		return changePageNumber(path, lastPageNumber)
	}
	return changePageNumber(path, page-1)
}

func changePageNumber(input string, to uint64) *string {
	re := regexp.MustCompile(`(page\[number\])(=\d+)`)
	strToParse := input
	if !re.MatchString(input) {
		strToParse = addParameterToPath(input, "page[number]=0")
	}
	result := re.ReplaceAllString(strToParse, fmt.Sprintf(`$1=%v`, to))
	return &result
}

func addParameterToPath(path, parameter string) string {
	var delimiter string
	if strings.Contains(path, "?") {
		delimiter = "&"
	} else {
		delimiter = "?"
	}
	return path + delimiter + parameter
}

func getLastPageNumber(limit, total uint64) uint64 {
	count := float64(total) / float64(limit)
	return uint64(math.Ceil(count))
}

package utils

import (
	"fmt"
	"strings"
)

func CombinSql(mainSql string, whereSql []string) string {
	if len(whereSql) < 1 {
		return strings.ReplaceAll(mainSql, "$WHERE", "")
	}
	str := " WHERE "
	for _, i := range whereSql {
		str += fmt.Sprintf(" %s ~", i)
	}
	str = strings.TrimRight(str, "~")
	str = strings.ReplaceAll(str, "~", "AND ")
	mainSql = strings.ReplaceAll(mainSql, "$WHERE", str)
	return mainSql
}

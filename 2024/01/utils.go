package main

import (
	"github.com/josephsilcock/adventofcode/2024/utils"
	"strings"
)

func getLineValues(line string) (ret [2]int) {
	parts := strings.Fields(line)

	ret[0] = utils.ConvertStringToInt(parts[0])
	ret[1] = utils.ConvertStringToInt(parts[1])

	return
}

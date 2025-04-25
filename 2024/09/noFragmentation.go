package main

import (
	"math"
)

func getChecksumWithNoFragmentation(fileMapSizes []int) (checksum int) {
	currentBlock := 0
	filesAdded := make(map[int]bool)

	for fileIndex, fileSize := range fileMapSizes {
		if math.Mod(float64(fileIndex), 2) == 0 && !filesAdded[fileIndex] {
			checksum += getChecksumContribution(fileIndex/2, currentBlock, fileSize)
			currentBlock += fileSize
			filesAdded[fileIndex] = true
		} else {
			for i := len(fileMapSizes) - 1; i > fileIndex; i -= 2 {
				reverseBlockFileSize := fileMapSizes[i]
				if filesAdded[i] {
					continue
				}
				if reverseBlockFileSize > fileSize {
					continue
				}
				checksum += getChecksumContribution(i/2, currentBlock, reverseBlockFileSize)
				currentBlock += reverseBlockFileSize
				fileSize -= reverseBlockFileSize
				filesAdded[i] = true
				if fileSize == 0 {
					break
				}
			}
			currentBlock += fileSize
		}
	}

	return checksum
}

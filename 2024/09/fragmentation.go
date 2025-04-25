package main

import "math"

func getChecksumWithFragmentation(fileMapSizes []int) (checksum int) {
	currentBlock := 0
	reverseIndex := len(fileMapSizes) - 1
	reverseBlockFileSizeLeft := fileMapSizes[reverseIndex]

FileLoop:
	for fileIndex, fileSize := range fileMapSizes {
		if math.Mod(float64(fileIndex), 2) == 0 {
			if fileIndex >= reverseIndex {
				break FileLoop
			}
			checksum += getChecksumContribution(fileIndex/2, currentBlock, fileSize)
			currentBlock += fileSize
		} else {
			for fileSize > 0 {
				if reverseIndex <= fileIndex {
					break FileLoop
				}
				reverseFileId := reverseIndex / 2
				if fileSize > reverseBlockFileSizeLeft {
					checksum += getChecksumContribution(reverseFileId, currentBlock, reverseBlockFileSizeLeft)
					currentBlock += reverseBlockFileSizeLeft
					fileSize -= reverseBlockFileSizeLeft
					reverseIndex -= 2
					reverseBlockFileSizeLeft = fileMapSizes[reverseIndex]
				} else {
					checksum += getChecksumContribution(reverseFileId, currentBlock, fileSize)
					currentBlock += fileSize
					reverseBlockFileSizeLeft -= fileSize
					fileSize = 0
				}
			}
		}
	}

	checksum += getChecksumContribution(reverseIndex/2, currentBlock, reverseBlockFileSizeLeft)

	return checksum
}

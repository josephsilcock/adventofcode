package main

func getChecksumContribution(fileId, startingBlock, fileSize int) (checksumContribution int) {
	for i := 0; i < fileSize; i++ {
		checksumContribution += fileId * (startingBlock + i)
	}
	return
}

package NFT_Server

type Blockchain_t struct {
	currHash uint8  // Previous hash
	size     uint   // Data Size
	data     string // Data
	nextHash uint8  // Next Hash
}

// checksum function
func checkSum(data map[int]uint8, size int) uint8 {
	var chkSum uint8
	for i := 0; i < size; i++ {
		chkSum += data[i]
	}
	return ^chkSum + 1
}

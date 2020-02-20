package e2e

// CheckSum8 Modulo 256.
func CrcMod256(data []byte) uint8 {
	var sum uint8
	for _, b := range data {
		sum += b
	}
	return sum & 0xff
}

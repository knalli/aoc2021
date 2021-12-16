package day16

var hex map[int32]string

func init() {
	hex = make(map[int32]string)
	hex['0'] = "0000"
	hex['1'] = "0001"
	hex['2'] = "0010"
	hex['3'] = "0011"
	hex['4'] = "0100"
	hex['5'] = "0101"
	hex['6'] = "0110"
	hex['7'] = "0111"
	hex['8'] = "1000"
	hex['9'] = "1001"
	hex['A'] = "1010"
	hex['B'] = "1011"
	hex['C'] = "1100"
	hex['D'] = "1101"
	hex['E'] = "1110"
	hex['F'] = "1111"
}

func decodeHex(input string) string {
	res := ""
	for _, c := range input {
		res += hex[c]
	}
	return res
}

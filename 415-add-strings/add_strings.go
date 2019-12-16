package add_strings

func addStrings(n1 string, n2 string) string {

	var res []byte
	if len(n1) > len(n2) {
		res = make([]byte, len(n1)+1)
	} else {
		res = make([]byte, len(n2)+1)
	}

	i := len(n1) - 1
	j := len(n2) - 1
	k := len(res) - 1
	carry := 0
	for {
		if i < 0 && j < 0 {
			break
		}

		var v1 int
		if i >= 0 {
			v1 = int(n1[i]) - '0'
		} else {
			v1 = 0
		}
		var v2 int
		if j >= 0 {
			v2 = int(n2[j]) - '0'
		} else {
			v2 = 0
		}

		n := v1 + v2 + carry
		if n <= 9 {
			carry = 0
		} else {
			carry = 1
			n -= 10
		}

		res[k] = byte(n + '0')
		i--
		j--
		k--
	}

	if carry == 1 {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])
}

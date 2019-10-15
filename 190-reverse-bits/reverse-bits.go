package reverse_bits

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return num
	}

	var res uint32

	shift := 0 // 偏移计数

	for {
		// 获取最低位
		if (num & 1) == 1 {
			res = res | 1 // res当前位赋值
		}

		num = num >> 1

		if num == 0 {
			break
		}

		res = res << 1
		shift++
	}

	if shift < 31 {
		res = res << uint32(31-shift)
	}

	return res
}

/**

解题思路

跳出条件：num不为0，进行操作

初始条件：
res = 0

b := num & 1：获得最低位的值
res | b： 得到当前位的值

b >> 1
res << 1


*/

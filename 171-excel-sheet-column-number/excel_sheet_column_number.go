package excel_sheet_column_number

func titleToNumber(s string) int {
	if s == "" {
		return 0
	}

	// base 表示基数，26^0 -> 26^1 -> 26^2
	base := 1

	total := 0

	for i := len(s) - 1; i >= 0; i-- {

		// 转换字符到数值
		val := int(s[i]-'A') + 1
		total += val * base
		base *= 26 // 基数进位
	}

	return total
}

package roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "测试1：正常顺序III",
			s:    "III",
			want: 3,
		},
		{
			name: "测试2：IV",
			s:    "IV",
			want: 4,
		},
		{
			"测试3：IX",
			"IX",
			9,
		},
		{
			"测试4：LVIII",
			"LVIII",
			58,
		},
		{
			"测试5：MCMXCIV",
			"MCMXCIV",
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

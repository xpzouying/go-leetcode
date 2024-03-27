package generateparenttheses

import (
	"encoding/json"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	s := generateParenthesis(1)

	data, _ := json.Marshal(s)
	t.Logf("generated: %s", data)

}

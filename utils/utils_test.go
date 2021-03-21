package utils

import "testing"

func TestTypeCheck(t *testing.T) {
	TypeCheck(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}

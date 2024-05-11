package bpn

import (
	"testing"
)

func TestMathFormulaToBPN(t *testing.T) {
	tests := []struct {
		param string
		want  string
		isErr bool
	}{
		{"5 + 3", "5 3 +", false},
		{"5 + 3 * 15", "5 3 15 * +", false},
		{"", "", true},
		{"5 + s", "", true},
		{"5 * 3 * 15", "5 3 15 * *", false},
		{"5 * 3 + 15", "5 3 * 15 +", false},
		{"5 + 3 + 15 / 6", "5 3 15 6 / + +", false},
		{"5 * 3 + 15 / 6", "5 3 * 15 6 / +", false},
		{"2 + 2 * 2", "2 2 2 * +", false},
		{"2 ^ 2", "2 2 ^", false},
		{"2 ^ 4", "2 4 ^", false},
		{"2 * 2 ^ 4", "2 2 4 ^ *", false},
		{"2 + 2 * 2 ^ 4", "2 2 2 4 ^ * +", false},
		{"( 2 + 2 ) * 2 ^ 4", "2 2 + 2 4 ^ *", false},
		{"2 + ( 2 * 2 ) ^ 4", "2 2 2 * 4 ^ +", false},
		{"2 + ( 1 * 2 + 3 ) ^ 4", "2 1 2 * 3 + 4 ^ +", false},
		{"2 * 2 + 3", "2 2 * 3 +", false},
		{"2 + ( 1 * ( 2 + 3 ) ) ^ 4", "2 1 2 3 + * 4 ^ +", false},
		{"2 + ( 1 * ( ( 2 + 3 ) ) ^ 4", "", true},
		{"2 + ( 1 * ( 2 + 3 ) ) ) ^ 4", "", true},
		{"2 + ( 1 * ( 2 + 3 ) + ) ^ 4", "", true},
		{"( 2 + 2 ) ^ 4 (", "", true},
		{"( 2 + 2 ) ^ 4", "2 2 + 4 ^", false},
		{"( 2 + 2 ) ^ 4 )", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.param, func(t *testing.T) {
			res, err := ToRPN(tt.param)

			if tt.isErr {
				if err == nil {
					t.Errorf("error must exists")
				} else if res != "" {
					t.Errorf("result must to be nil. res: %+v", res)
				}
			} else {
				if err != nil {
					t.Errorf("error doesn't must exist. error %+v", err)
				} else if res != tt.want {
					t.Errorf("result %+v not equal to expected %+v", res, tt.want)
				}
			}
		})
	}
}

func TestCalcBPN(t *testing.T) {
	tests := []struct {
		name  string
		param string
		want  float64
		isErr bool
	}{
		{"5 + 3", "5 3 +", 8, false},
		{"5 + 3 * 15", "5 3 15 * +", 50, false},
		{"", "", 0, true},
		{"5 + s", "5 s +", 0, true},
		{"5+s", "5s +", 0, true},
		{"5s", "5s", 0, true},
		{"5 * 3 * 15", "5 3 15 * *", 225, false},
		{"5 * 3 + 15", "5 3 * 15 +", 30, false},
		{"5 + 3 + 15 / 6", "5 3 15 6 / + +", 10.5, false},
		{"5 * 3 + 15 / 6", "5 3 * 15 6 / +", 17.5, false},
		{"2 + 2 * 2", "2 2 2 * +", 6, false},
		{"2 ^ 2", "2 2 ^", 4, false},
		{"2 ^ 4", "2 4 ^", 16, false},
		{"2 * 2 ^ 3", "2 2 3 ^ *", 16, false},
		{"2 * 2 ^ 4", "2 2 4 ^ *", 32, false},
		{"2 + 2 * 2 ^ 4", "2 2 2 4 ^ * +", 34, false},
		{"5 - 3 * 15", "5 3 15 * -", -40, false},
		{"2 + ( 1 * ( 2 + 3 ) ) ^ 4", "2 1 2 3 + * 4 ^ +", 627, false},
		{"( 2 + 2 ) ^ 4", "2 2 + 4 ^", 256, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CalcRPN(tt.param)

			if tt.isErr {
				if err == nil {
					t.Errorf("error must exists")
				} else if res != 0 {
					t.Errorf("result %+v expected 0", res)
				}
			}
			if res != tt.want {
				t.Errorf("result %+v is not expected %+v", res, tt.want)
			} else if !tt.isErr && err != nil {
				t.Errorf("error %+v isn't nil", err)
			}
		})
	}
}

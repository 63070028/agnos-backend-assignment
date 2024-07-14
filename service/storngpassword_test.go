package service

import (
	"testing"

	"github.com/63070028/agnos-backend-assignment/model"
)

func TestMiminimumActions(t *testing.T) {

	type args struct {
		password string
		config   model.ConfigStrongPassword
	}

	config := model.ConfigStrongPassword{
		MinLowerCase: 1,
		MinUpperCase: 1,
		MinDigit:     1,
		Repeat:       3,
		MinLength:    6,
		MaxLength:    19,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"test case 1", args{".", config}, 5},
		{"test case 2", args{"a", config}, 5},
		{"test case 3", args{"aA1", config}, 3},
		{"test case 4", args{"1445D1cd", config}, 0},
		{"test case 5", args{".....", config}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiminimumActions(tt.args.password, tt.args.config); got != tt.want {
				t.Errorf("MiminimumActions() = %v, want %v", got, tt.want)
			}
		})
	}
}

package parser

import (
	"reflect"
	"testing"
)

func TestNewSimpleParser(t *testing.T) {
	tests := []struct {
		name string
		want *SimpleParser
	}{
		{
			"initSimpleParser",
			&SimpleParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSimpleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimpleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleParser_ParseAndSum(t *testing.T) {
	type args struct {
		numbers string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"testReturnsZeroWhenEmptyString",
			args{numbers: ""},
			0,
			false,
		},
		{
			"testReturnErrInvalidParseWhenWrongFormat",
			args{numbers: "1 2"},
			0,
			true,
		},
		{
			"testReturnErrInvalidParseWhenParseCharacter",
			args{numbers: "s"},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleParser{}
			got, err := s.ParseAndSum(tt.args.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAndSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAndSum() got = %v, want %v", got, tt.want)
			}
		})
	}
}

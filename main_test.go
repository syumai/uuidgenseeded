package main

import (
	"io"
	"strings"
	"testing"
)

func Test_realMain(t *testing.T) {
	tests := map[string]struct {
		seedReader io.Reader
		lower      bool
		want    string
	}{
		"ok with upper case": {
			seedReader: strings.NewReader("example"),
			lower: false,
			want: "475A87DF-68E7-4C04-A4F0-E11504AEC553",
		},
		"ok with lower case": {
			seedReader: strings.NewReader("example"),
			lower: true,
			want: "475a87df-68e7-4c04-a4f0-e11504aec553",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := realMain(tt.seedReader, tt.lower)
			if err != nil {
				t.Errorf("realMain() error = %v, wantErr %v", err, false)
				return
			}
			if got != tt.want {
				t.Errorf("realMain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

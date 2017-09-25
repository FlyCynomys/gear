package format

import (
	"testing"
	"time"
)

func TestToInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"p2p", args{"1"}, 1},
		{"0p2p", args{"01"}, 1},
		{"020", args{"0"}, 0},
		{"0020", args{"00"}, 0},
		{"n2n", args{"-1"}, -1},
		{"0n2n", args{"-01"}, -1},
		{"errfloat2int", args{"1.1"}, 0},
		{"errstr2int", args{"hello"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.v); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"p2p", args{"1"}, 1},
		{"0p2p", args{"01"}, 1},
		{"020", args{"0"}, 0},
		{"0020", args{"00"}, 0},
		{"n2n", args{"-1"}, -1},
		{"0n2n", args{"-01"}, -1},
		{"errfloat2int", args{"1.1"}, 0},
		{"errstr2int", args{"hello"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.v); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeFromTi(t *testing.T) {
	type args struct {
		template  string
		timepoint time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"dateLocal", args{"2006-01-02", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local)}, "2017-05-21"},
		{"dateUTC", args{"2006-01-02", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.UTC)}, "2017-05-21"},
		{"datetimeUTC", args{"2006-01-02 15:04:05", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.UTC)}, "2017-05-21 14:00:00"},
		{"datetimeLocal", args{"2006-01-02 15:04:05", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local)}, "2017-05-21 14:00:00"},
		{"dateSlash", args{"2006/01/02", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local)}, "2017/05/21"},
		{"datetimeSlash", args{"2006/01/02 15/04/05", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local)}, "2017/05/21 14/00/00"},
		{"dateFormat", args{"2006-01-02", time.Date(2006, time.January, 02, 14, 0, 0, 0, time.Local)}, "2006-01-02"},
		{"datetimeFormat", args{"2006-01-02 15:04:05", time.Date(2006, time.January, 02, 15, 04, 05, 0, time.Local)}, "2006-01-02 15:04:05"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeFromTi(tt.args.template, tt.args.timepoint); got != tt.want {
				t.Errorf("FormatTimeFromTi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeFromUnix(t *testing.T) {
	type args struct {
		template  string
		timepoint int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"dateLocal", args{"2006-01-02", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local).Unix()}, "2017-05-21"},
		{"dateUTC", args{"2006-01-02", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.UTC).Unix()}, "2017-05-21"},
		{"datetimeUTC", args{"2006-01-02 15:04:05", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.UTC).Unix()}, "2017-05-21 14:00:00"},
		{"datetimeLocal", args{"2006-01-02 15:04:05", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local).Unix()}, "2017-05-21 14:00:00"},
		{"dateSlash", args{"2006/01/02", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local).Unix()}, "2017/05/21"},
		{"datetimeSlash", args{"2006/01/02 15/04/05", time.Date(2017, time.May, 21, 14, 0, 0, 0, time.Local).Unix()}, "2017/05/21 14/00/00"},
		{"dateFormat", args{"2006-01-02", time.Date(2006, time.January, 02, 14, 0, 0, 0, time.Local).Unix()}, "2006-01-02"},
		{"datetimeFormat", args{"2006-01-02 15:04:05", time.Date(2006, time.January, 02, 15, 04, 05, 0, time.Local).Unix()}, "2006-01-02 15:04:05"},
		{"datetimeFormatUTC", args{"2006-01-02 15:04:05", time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC).Unix()}, "2006-01-02 15:04:05"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeFromUnix(tt.args.template, tt.args.timepoint); got != tt.want {
				t.Errorf("FormatTimeFromUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToUnix(t *testing.T) {
	type args struct {
		timepoint string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"testLocal", args{"2017-05-21 05:20:00"}, time.Date(2017, time.May, 21, 05, 20, 0, 0, time.Local).Unix()},
		{"testUTC", args{"2017-05-21 05:20:00"}, time.Date(2017, time.May, 21, 05, 20, 0, 0, time.UTC).Unix()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeToUnix(tt.args.timepoint); got != tt.want {
				t.Errorf("FormatTimeToUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

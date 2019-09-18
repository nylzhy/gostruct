package gostruct

import (
	"reflect"
	"testing"
)

func TestCalcSizeof(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantN int
	}{
		{
			name:  "normal1",
			args:  args{"cbB?Hi"},
			wantN: 10,
		},
		{
			name:  "normal2",
			args:  args{"dfQ?Hi"},
			wantN: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := CalcSizeof(tt.args.s); gotN != tt.wantN {
				t.Errorf("CalcSizeof() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestBoolsToBytes(t *testing.T) {
	type args struct {
		b  []bool
		et EndianType
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
		wantErr bool
	}{
		{
			name:    "bool_1",
			args:    args{[]bool{true, false, false, true, false, true, false, false}, BigEndian},
			wantRes: []byte{41},
			wantErr: false,
		},
		{
			name:    "bool_2",
			args:    args{[]bool{true, false, false, true, false, true}, BigEndian},
			wantRes: []byte{41},
			wantErr: false,
		},
		{
			name:    "bool_3",
			args:    args{[]bool{true, true, true, true, true, true, true}, BigEndian},
			wantRes: []byte{127},
			wantErr: false,
		},
		{
			name:    "bool_4",
			args:    args{[]bool{true}, BigEndian},
			wantRes: []byte{1},
			wantErr: false,
		},
		{
			name:    "bool_5",
			args:    args{[]bool{true, false, false, true, true, false, false, true, false, true, false, false, true}, LittleEndian},
			wantRes: []byte{153, 18},
			wantErr: false,
		},
		{
			name:    "bool_5_2",
			args:    args{[]bool{true, false, false, true, true, false, false, true, false, true, false, false, true}, BigEndian},
			wantRes: []byte{18, 153},
			wantErr: false,
		},
		{
			name:    "bool_6",
			args:    args{[]bool{true, false, false, true, true, false, false, true}, LittleEndian},
			wantRes: []byte{153},
			wantErr: false,
		},
		{
			name:    "bool_7",
			args:    args{[]bool{false, true, false, false, true}, LittleEndian},
			wantRes: []byte{18},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := BoolsToBytes(tt.args.b, tt.args.et)
			if (err != nil) != tt.wantErr {
				t.Errorf("BoolsToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BoolsToBytes() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestBytesToBools(t *testing.T) {
	type args struct {
		b  []byte
		et EndianType
	}
	tests := []struct {
		name    string
		args    args
		wantRes []bool
		wantErr bool
	}{
		{
			name:    "TEST1",
			args:    args{[]byte{0x34}, BigEndian},
			wantRes: []bool{false, false, true, false, true, true, false, false},
			wantErr: false,
		},
		{
			name:    "TEST2",
			args:    args{[]byte{0x34, 0x0d}, BigEndian},
			wantRes: []bool{true, false, true, true, false, false, false, false, false, false, true, false, true, true, false, false},
			wantErr: false,
		},
		{
			name:    "TEST3",
			args:    args{[]byte{0x0d, 0x34}, LittleEndian},
			wantRes: []bool{true, false, true, true, false, false, false, false, false, false, true, false, true, true, false, false},
			wantErr: false,
		},
		{
			name:    "TEST4",
			args:    args{[]byte{0x0d, 0x34}, BigEndian},
			wantRes: []bool{false, false, true, false, true, true, false, false, true, false, true, true, false, false, false, false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := BytesToBools(tt.args.b, tt.args.et)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToBools() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BytesToBools() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

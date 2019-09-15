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

func TestInt16ToBytes(t *testing.T) {
	type args struct {
		n  int16
		et EndianType
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
		wantErr bool
	}{
		{
			name:    "int16tobytes",
			args:    args{-128, LittleEndian},
			wantRes: []byte{0x80, 0xFF},
			wantErr: false,
		},
		{
			name:    "int16tobytes",
			args:    args{-128, BigEndian},
			wantRes: []byte{0xFF, 0x80},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Int16ToBytes(tt.args.n, tt.args.et)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int16Tobytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Int16Tobytes() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

package gostruct

import "testing"

func Test_transview(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantR   string
		wantErr bool
	}{
		{
			name:    "normal1",
			args:    args{"i2b"},
			wantR:   "ibb",
			wantErr: false,
		},
		{
			name:    "normal2",
			args:    args{"2i10c6b"},
			wantR:   "iiccccccccccbbbbbb",
			wantErr: false,
		},
		{
			name:    "normal2",
			args:    args{"0i10c06b"},
			wantR:   "ccccccccccbbbbbb",
			wantErr: false,
		},
		{
			name:    "err1",
			args:    args{"0i10A06b"},
			wantR:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := transview(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("transview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("transview() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

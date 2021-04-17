package main

import "testing"

func TestXXXReader_Read(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		r       XXXReader
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := XXXReader{}
			gotN, err := r.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("XXXReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("XXXReader.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestXXXReader_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		r       XXXReader
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := XXXReader{}
			gotN, err := r.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("XXXReader.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("XXXReader.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestXXXReader_Close(t *testing.T) {
	tests := []struct {
		name    string
		r       XXXReader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := XXXReader{}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("XXXReader.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

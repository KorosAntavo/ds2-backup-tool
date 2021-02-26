package loop

import "testing"

// not a unit test
func Test_userHomeDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test0",
			// not a unit test
			want: userHomeDir(),
		},
	}

	var got string

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got = userHomeDir(); got != tt.want {
				t.Errorf("userHomeDir() = %v, want %v", got, tt.want)
			}
		})

		t = t
	}
}

func Test_performAction(t *testing.T) {
	type args struct {
		index int
		down  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test0 - save",
			args: struct {
				index int
				down  bool
			}{index: 0, down: false},
			wantErr: false,
		},
		{
			name: "test0 - load",
			args: struct {
				index int
				down  bool
			}{index: 0, down: true},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := performAction(tt.args.index, tt.args.down); (err != nil) != tt.wantErr {
				t.Errorf("performAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

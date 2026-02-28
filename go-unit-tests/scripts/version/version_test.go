package version

import "testing"

func Test_GoMinorVersion(t *testing.T) {
	type args struct {
		versionOutput []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				versionOutput: []byte("go version go1.19.2 linux/amd64"),
			},
			want: 19,
		},
		{
			name: "too long patch version with some letters at beginning",
			args: args{
				versionOutput: []byte("go version go1.19.gertgergrege linux/amd64"),
			},
			want: 19,
		},
		{
			name: "patch version missing",
			args: args{
				versionOutput: []byte("go version go1.16 linux/amd64"),
			},
			want: 16,
		},
		{
			name: "some stupid quoute with go version smh",
			args: args{
				versionOutput: []byte(
					"The problem is not that I use metaclasses. It's that I go1.19.2 gleefully use metaclasses.",
				),
			},
			want: 19,
		},
		{
			name: "HURRAY! GO 123456 released with support for quantum computing",
			args: args{
				versionOutput: []byte("go version go123456.19.2"),
			},
			want: 19,
		},
		{
			name: "wrong Go version without minor part",
			args: args{
				versionOutput: []byte("go version 2"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GoMinorVersion(tt.args.versionOutput)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoMinorVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GoMinorVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
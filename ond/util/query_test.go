package util

import "testing"

type Input struct {
	Name    string `url:"name"`
	Age     string `url:"age" json:"age"`
	Address string `json:"address"`
}

var (
	input1 = Input{
		Name: "tom",
		Age:  "3",
	}
	input2 = Input{
		Name: "tom",
	}
	input3 = Input{
		Name: "",
		Age:  "3",
	}
)

func TestBuildQuery(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				data: input1,
			},
			want:    "age=3&name=tom",
			wantErr: false,
		},
		{
			name: "test 2",
			args: args{
				data: input2,
			},
			want:    "name=tom",
			wantErr: false,
		},
		{
			name: "test 3",
			args: args{
				data: input3,
			},
			want:    "age=3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildQuery(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

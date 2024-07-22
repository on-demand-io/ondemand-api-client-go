package util

import "testing"

type Input struct {
	Name    string `url:"name"`
	Age     string `url:"age" json:"age"`
	Address string `json:"address"`
}

type InputWithSlice struct {
	Plugins []string `url:"pluginId"`
	Age     int      `url:"age"`
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
	input4 = InputWithSlice{
		Plugins: []string{"p1", "p2"},
		Age:     2,
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
		{
			name: "test 4 - slice",
			args: args{
				data: input4,
			},
			want:    "age=2&pluginId=p1&pluginId=p2",
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

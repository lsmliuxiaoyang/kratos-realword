package data

import (
	"kratos-realwd/internal/conf"
	"testing"
)

func TestNewDb(t *testing.T) {
	type args struct {
		c *conf.Data
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				c: &conf.Data{
					Database: &conf.Data_Database{
						Driver: "mysql",
						Source: "root:123456@tcp(127.0.0.1:3306)/testDb?charset=utf8&parseTime=true&loc=Local",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewDb(tt.args.c)
		})
	}
}

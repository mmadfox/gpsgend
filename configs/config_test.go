package config

import (
	"reflect"
	"testing"
)

func TestFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    func() *Config
		wantErr bool
	}{
		{
			name:    "should return error when file does not exist",
			args:    args{filename: "somefile"},
			wantErr: true,
		},
		{
			name:    "should return error when invalid file format",
			args:    args{filename: "./testdata/badconfig"},
			wantErr: true,
		},
		{
			name: "should return config",
			args: args{filename: "./testdata/config.yaml"},
			want: func() *Config {
				conf := new(Config)
				conf.Service = "gpsgend"
				conf.Logger.Format = "text"
				conf.Logger.Level = "info"
				conf.API.HTTP.Addr = "0.0.0.0:1"
				conf.Storage.Mongodb.CollectionName = "devices"
				conf.Storage.Mongodb.DatabaseName = "gpsgend"
				conf.Storage.Mongodb.URI = "mongodb://127.0.0.1:27017/gpsgend"
				return conf
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("FromFile() = %v, want %v", got, want)
			}
		})
	}
}

package http

import (
	"reflect"
	"testing"

	"github.com/mmadfox/gpsgend/internal/device"
)

func Test_decodeDeviceQueryFilter(t *testing.T) {
	str2ptr := func(s string) *string {
		return &s
	}
	type args struct {
		c func() queryer
	}
	tests := []struct {
		name    string
		args    args
		wantQf  device.QueryFilter
		wantErr bool
	}{
		{
			name: "should return nil when model param not set",
			args: args{
				c: func() queryer { return newQueryMock() },
			},
			wantQf:  device.QueryFilter{},
			wantErr: false,
		},
		{
			name: "should return nil when model param set and is empty value",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("model", " ")
					return q
				},
			},
			wantQf:  device.QueryFilter{},
			wantErr: false,
		},
		{
			name: "should return string ptr when model param set and is not empty",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("model", "TestModel")
					return q
				},
			},
			wantQf: device.QueryFilter{
				Model: str2ptr("TestModel"),
			},
			wantErr: false,
		},
		{
			name: "should return str slice when device param set with comma delimiter",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("device", "id1,id2")
					return q
				},
			},
			wantQf: device.QueryFilter{
				ID: &[]string{"id1", "id2"},
			},
			wantErr: false,
		},
		{
			name: "should return nil when device param set and is empty value",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("device", " ")
					return q
				},
			},
			wantQf:  device.QueryFilter{},
			wantErr: false,
		},
		{
			name: "should return nil when device param only comma",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("device", ",")
					return q
				},
			},
			wantQf:  device.QueryFilter{},
			wantErr: false,
		},
		{
			name: "should return status param",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("status", "1,2")
					return q
				},
			},
			wantQf: device.QueryFilter{
				Status: &[]int{1, 2},
			},
			wantErr: false,
		},
		{
			name: "should return sensor param",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("sensor", "1,2")
					return q
				},
			},
			wantQf: device.QueryFilter{
				Sensor: &[]string{"1", "2"},
			},
			wantErr: false,
		},
		{
			name: "should return user param",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("user", "1,2")
					return q
				},
			},
			wantQf: device.QueryFilter{
				User: &[]string{"1", "2"},
			},
			wantErr: false,
		},
		{
			name: "should return error when limit param is invalid",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("limit", ">")
					return q
				},
			},
			wantErr: true,
		},
		{
			name: "should return nil when limit param is negative",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("limit", "-1")
					return q
				},
			},
			wantQf: device.QueryFilter{},
		},
		{
			name: "should return limit param",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("limit", "100")
					return q
				},
			},
			wantQf: device.QueryFilter{
				Limit: 100,
			},
		},
		{
			name: "should return error when page param is invalid",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("page", ">")
					return q
				},
			},
			wantErr: true,
		},
		{
			name: "should return nil when page param is negative",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("page", "-1")
					return q
				},
			},
			wantQf: device.QueryFilter{},
		},
		{
			name: "should return page param",
			args: args{
				c: func() queryer {
					q := newQueryMock()
					q.set("page", "1")
					return q
				},
			},
			wantQf: device.QueryFilter{
				Page: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQf, err := decodeDeviceQueryFilter(tt.args.c())
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeDeviceQueryFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotQf, tt.wantQf) {
				t.Errorf("decodeDeviceQueryFilter() = %v, want %v", gotQf, tt.wantQf)
			}
		})
	}
}

type queryMock struct {
	params map[string]string
}

func (qm queryMock) Query(key string, defaultValue ...string) string {
	val, ok := qm.params[key]
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
	}
	return val
}

func (qm queryMock) set(key string, value string) {
	qm.params[key] = value
}

func newQueryMock() queryMock {
	return queryMock{params: make(map[string]string)}
}

package tsl

import (
	"reflect"
	"testing"
	"time"
)

func TestQuery_Select(t *testing.T) {
	type fields struct {
		raw string
	}
	type args struct {
		metric string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{{
		name: "plain query",
		fields: fields{
			raw: "",
		},
		args: args{
			metric: "os.cpu",
		},
		want: &Query{
			raw: "select(\"os.cpu\")",
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				raw: tt.fields.raw,
			}
			got := q.Select(tt.args.metric)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Last(t *testing.T) {
	type fields struct {
		raw string
	}
	type args struct {
		d  time.Duration
		at time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{{
		name: "Last 5 minutes",
		fields: fields{
			raw: "",
		},
		args: args{
			d:  5 * time.Minute,
			at: NilTime,
		},
		want: &Query{
			raw: ".last(5m)",
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				raw: tt.fields.raw,
			}

			got := q.Last(tt.args.d, tt.args.at)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

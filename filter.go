package hmgdb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 使用示例： NewF().Str("name",name).Str("k",v)
func NewF() Filter {
	return Filter{}
}

type Filter bson.M

func (f Filter) Str(k string, v string) Filter {
	if v != "" {
		f[k] = v
	}
	return f
}

func (f Filter) Int(k string, v int) Filter {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f Filter) Int32(k string, v int32) Filter {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f Filter) Int64(k string, v int64) Filter {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f Filter) Time(k string, v time.Time) Filter {
	if !v.IsZero() {
		f[k] = v
	}
	return f
}

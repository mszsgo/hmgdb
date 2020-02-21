package hmgdb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 使用示例： NewF().Str("name",name).Str("k",v)
// Deprecated: 改为使用 hmgdb.M{}.Str(k,v) ,
func NewF() M {
	return M{}
}

type M bson.M

func (f M) M(k string, v M) M {
	if len(v) > 0 {
		f[k] = v
	}
	return f
}

func (f M) Str(k string, v string) M {
	if v != "" {
		f[k] = v
	}
	return f
}

func (f M) Int(k string, v int) M {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f M) Int32(k string, v int32) M {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f M) Int64(k string, v int64) M {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f M) Time(k string, v time.Time) M {
	if !v.IsZero() {
		f[k] = v
	}
	return f
}

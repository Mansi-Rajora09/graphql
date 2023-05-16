package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func AutoField(k string, v interface{}) zapcore.Field {
	switch v := v.(type) {
	case string:
		return zap.String(k, v)
	case *string:
		return zap.String(k, *v)
	case int64:
		return zap.Int64(k, v)
	case *int64:
		return zap.Int64(k, *v)
	case int:
		return zap.Int(k, v)
	case float64:
		return zap.Float64(k, v)
	case bool:
		return zap.Bool(k, v)
	case time.Time:
		return zap.Time(k, v)
	default:
		if e, ok := v.(error); ok {
			return zap.String(k, e.Error())
		} else if v == nil {
			return zap.String(k, "nil")
		} else {
			return zap.Any(k, v)
		}
	}
}

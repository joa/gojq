package gojq

import (
	"encoding/json"
	"math"
	"math/big"
	"strings"
)

func normalizeNumber(v json.Number) any {
	if i, err := v.Int64(); err == nil && math.MinInt <= i && i <= math.MaxInt {
		return int(i)
	}
	if strings.ContainsAny(v.String(), ".eE") {
		if f, err := v.Float64(); err == nil {
			return f
		}
	}
	if bi, ok := new(big.Int).SetString(v.String(), 10); ok {
		return bi
	}
	if strings.HasPrefix(v.String(), "-") {
		return math.Inf(-1)
	}
	return math.Inf(1)
}

// Normalize any given data.
//
// This function will in particular change any given map[any]any
// so that all the integer values will be of type int and all the
// floating point values will be of type float64. This method will
// modify the given data in place.
//
// Typically running a gojq query will normalize the data automatically
// unless the [WithoutNormalize] option has been specified.
func Normalize(v any) any {
	switch v := v.(type) {
	case json.Number:
		return normalizeNumber(v)
	case *big.Int:
		if v.IsInt64() {
			if i := v.Int64(); math.MinInt <= i && i <= math.MaxInt {
				return int(i)
			}
		}
		return v
	case int64:
		if math.MinInt <= v && v <= math.MaxInt {
			return int(v)
		}
		return big.NewInt(v)
	case int32:
		return int(v)
	case int16:
		return int(v)
	case int8:
		return int(v)
	case uint:
		if v <= math.MaxInt {
			return int(v)
		}
		return new(big.Int).SetUint64(uint64(v))
	case uint64:
		if v <= math.MaxInt {
			return int(v)
		}
		return new(big.Int).SetUint64(v)
	case uint32:
		if uint64(v) <= math.MaxInt {
			return int(v)
		}
		return new(big.Int).SetUint64(uint64(v))
	case uint16:
		return int(v)
	case uint8:
		return int(v)
	case float32:
		return float64(v)
	case []any:
		for i, x := range v {
			v[i] = Normalize(x)
		}
		return v
	case map[string]any:
		for k, x := range v {
			v[k] = Normalize(x)
		}
		return v
	default:
		return v
	}
}

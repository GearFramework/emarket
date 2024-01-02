package gear

import "strconv"

func AtoUI8(value string, def uint8) uint8 {
	val := uint8(0)
	if i, err := strconv.ParseInt(value, 10, 8); err == nil {
		val = uint8(i)
	}
	return val
}

func AtoUI(value string, def uint) uint {
	val := def
	if i, err := strconv.Atoi(value); err == nil {
		val = uint(i)
	}
	return val
}

func AtoUI16(value string, def uint16) uint16 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 16); err == nil {
		val = uint16(i)
	}
	return val
}

func AtoUI32(value string, def uint32) uint32 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 32); err == nil {
		val = uint32(i)
	}
	return val
}

func AtoUI64(value string, def uint64) uint64 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		val = uint64(i)
	}
	return val
}
func AtoI8(value string, def int8) int8 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 8); err == nil {
		val = int8(i)
	}
	return val
}

func AtoI(value string, def int) int {
	val := def
	if i, err := strconv.Atoi(value); err == nil {
		val = int(i)
	}
	return val
}

func AtoI16(value string, def int16) int16 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 16); err == nil {
		val = int16(i)
	}
	return val
}

func AtoI32(value string, def int32) int32 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 32); err == nil {
		val = int32(i)
	}
	return val
}

func AtoI64(value string, def int64) int64 {
	val := def
	if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		val = int64(i)
	}
	return val
}

func UI8ToA(value uint8) string {
	return strconv.FormatUint(uint64(value), 10)
}

func UIToA(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

func UI16ToA(value uint16) string {
	return strconv.FormatUint(uint64(value), 10)
}

func UI32ToA(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

func UI64ToA(value uint64) string {
	return strconv.FormatUint(uint64(value), 10)
}

func I8ToA(value int8) string {
	return strconv.FormatInt(int64(value), 10)
}

func IToA(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func I16ToA(value int16) string {
	return strconv.FormatInt(int64(value), 10)
}

func I32ToA(value int32) string {
	return strconv.FormatInt(int64(value), 10)
}

func I64ToA(value int64) string {
	return strconv.FormatInt(int64(value), 10)
}

func GetMapKeys(mapa map[string]any) []string {
	r := []string{}
	for k := range mapa {
		r = append(r, k)
	}
	return r
}

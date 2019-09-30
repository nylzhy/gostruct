package gostruct

import (
	"fmt"
	"strconv"
	"strings"
)

type encoder struct {
	endian    EndianType
	Formatter string
}

//New return a struct object
func New(f string, et EndianType) (Encoder, error) {
	d := &encoder{
		endian: et,
	}
	ft, err := transview(f)
	if err != nil {
		return nil, err
	}
	d.Formatter = ft
	return d, nil
}

func (e *encoder) Unpack(b []byte) (data []interface{}, err error) {
	if len(b) < e.calcsize() {
		return nil, fmt.Errorf("given data length %d do not match the 'encoder' type length %d", len(b), e.calcsize())
	}
	il := 0
	ir := 0
	for _, v := range e.Formatter {
		switch string(v) {
		case "c":
			ir = il + 1
			data = append(data, BytesToChar(b[il:ir]))
		case "b":
			ir = il + 1
			temp, err := BytesToInt8(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "B":
			ir = il + 1
			temp, err := BytesToUint8(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "h":
			ir = il + 2
			temp, err := BytesToInt16(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "H":
			ir = il + 2
			temp, err := BytesToUint16(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "i", "l":
			ir = il + 4
			temp, err := BytesToInt32(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "I", "L":
			ir = il + 4
			temp, err := BytesToUint32(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "q":
			ir = il + 8
			temp, err := BytesToInt64(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "Q":
			ir = il + 8
			temp, err := BytesToUint64(b[il:ir], e.endian)
			if err != nil {
				data = append(data, 0)
			}
			data = append(data, temp)
		case "f":
			ir = il + 4
			temp, err := BytesToFloat(b[il:ir], e.endian)
			if err != nil {
				data = append(data, -1.0)
			}
			data = append(data, temp)
		case "d":
			ir = il + 8
			temp, err := BytesToDouble(b[il:ir], e.endian)
			if err != nil {
				data = append(data, -1.0)
			}
			data = append(data, temp)
		case "?":
			ir = il + 1
			temp, err := BytesToBools(b[il:ir], e.endian)
			if err != nil {
				data = append(data, nil)
			}
			data = append(data, temp)
		}
		il = ir
	}
	return
}

func (e *encoder) Pack(data []interface{}) (b []byte, err error) {
	for i, v := range e.Formatter {
		switch string(v) {
		case "c":
			b = append(b, CharToBytes(data[i].(string))...)
		case "b":
			temp, err := Int8ToBytes(data[i].(int8), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)

		case "B":
			temp, err := Uint8ToBytes(data[i].(uint8), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "h":
			temp, err := Int16ToBytes(data[i].(int16), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "H":
			temp, err := Uint16ToBytes(data[i].(uint16), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "i", "l":
			temp, err := Int32ToBytes(data[i].(int32), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "I", "L":
			temp, err := Uint32ToBytes(data[i].(uint32), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "q":
			temp, err := Int64ToBytes(data[i].(int64), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "Q":
			temp, err := Uint64ToBytes(data[i].(uint64), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "f":
			temp, err := FloatToBytes(data[i].(float32), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "d":
			temp, err := DoubleToBytes(data[i].(float64), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		case "?":
			temp, err := BoolsToBytes(data[i].([]bool), e.endian)
			if err != nil {
				return b, nil
			}
			b = append(b, temp...)
		}
	}
	return
}

func (e *encoder) calcsize() (n int) {
	dtysize := map[string]int{"c": 1, "b": 1, "B": 1, "?": 1, "h": 2, "H": 2, "i": 4, "I": 4, "l": 4, "L": 4, "q": 8, "Q": 8, "f": 4, "d": 8}
	for _, v := range e.Formatter {
		n += dtysize[string(v)]
	}
	return
}

func transview(s string) (r string, err error) {
	dtysize := map[string]int{"c": 1, "b": 1, "B": 1, "?": 1, "h": 2, "H": 2, "i": 4, "I": 4, "l": 4, "L": 4, "q": 8, "Q": 8, "f": 4, "d": 8}
	//标识数字计数是否开始
	var counting bool
	var numberstr string
	for _, v := range s {
		//判断是否想要的字符
		if _, ok := dtysize[string(v)]; ok {
			//如果前面又数字为处理，则重复字符
			if counting {
				n, err := strconv.ParseUint(numberstr, 10, 32)
				if err != nil {
					return "", err
				}
				r += strings.Repeat(string(v), int(n))
				//重复完成后重置计数状态
				counting = false
			} else {
				//如果没有数字要处理，则添加
				r += string(v)
			}
		} else if 48 <= v && v < 58 {
			//如果不是格式化字符，则必须是0-9的数字，支持0将字符删除
			if !counting {
				numberstr = string(v)
				counting = true
			} else {
				numberstr += string(v)
			}
		} else {
			//其他非数字或字符的字符将返回错误
			return "", fmt.Errorf("%v in format string %vcan't be recognized", string(v), s)
		}
	}
	return
}

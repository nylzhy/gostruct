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
	for _, v := range e.Formatter {
		switch string(v) {
		case "c":
			data = append(data)
		case "b":
			return
		case "B":
			return
		}
	}
	return
}

func (e *encoder) Pack(data []interface{}) (b []byte, err error) {
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

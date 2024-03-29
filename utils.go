package gostruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//CalcSizeof calculate []byte lenght of the encoder format string
func CalcSizeof(s string) (n int) {
	dtysize := map[string]int{"c": 1, "b": 1, "B": 1, "?": 1, "h": 2, "H": 2, "i": 4, "I": 4, "l": 4, "L": 4, "q": 8, "Q": 8, "f": 4, "d": 8}
	for _, v := range s {
		n += dtysize[string(v)]
	}
	return
}

// CharToBytes convert string to byte, code 'c'
func CharToBytes(c string) (res []byte) {
	return []byte(c)
}

// BytesToChar convert bytes to string, code 'c'
func BytesToChar(b []byte) (c string) {
	return string(b)
}

// ByteToBools convert byte to bools, code '?'
func ByteToBools(b byte) (res []bool) {
	dict := map[string]bool{"0": false, "1": true}
	result := make([]bool, 8)
	bstr := fmt.Sprintf("%b", b)
	for i := range bstr {
		result[len(bstr)-i-1] = dict[string(bstr[i])]
	}
	return result
}

// BytesToBools convert bytes to bools, code '?'
func BytesToBools(b []byte, et EndianType) (res []bool, err error) {
	for i := 0; i < len(b); i++ {
		if et {
			temp := ByteToBools(b[len(b)-i-1])
			res = append(res, temp...)
			continue
		}
		temp := ByteToBools(b[i])
		res = append(res, temp...)
	}
	return
}

// BoolsToByte convert bools to byte, code '?' convert less 8 bool data to bytes
func BoolsToByte(b []bool) (res []byte, err error) {
	if len(b) > 8 {
		return []byte{0}, fmt.Errorf("input bool beyong the byte cap")
	}
	var n byte
	for i := range b {
		n = n << 1
		if b[len(b)-i-1] {
			n = n + 1
		}
	}
	return []byte{n}, nil
}

// BoolsToBytes convert bools to byte, code '?' convert many bools data to bytes
func BoolsToBytes(b []bool, et EndianType) (res []byte, err error) {
	bsize := len(b)
	if bsize <= 8 {
		return BoolsToByte(b)
	}
	resulte := []byte{}
	for i := 0; i < bsize/8+1; i++ {
		left := i * 8
		right := i*8 + 8
		if right > bsize {
			right = bsize
		}
		temp, _ := BoolsToByte(b[left:right])
		resulte = append(resulte, temp...)
	}
	if et {
		for i := range resulte {
			res = append(res, resulte[len(resulte)-i-1])
		}
	} else {
		res = resulte
	}
	return res, nil
}

// BoolToByte convert bool to byte, code '?' for one bool convert []byte{} with 1 length
func BoolToByte(b bool) (res []byte, err error) {
	if b {
		return []byte{0x01}, nil
	}
	return []byte{0x00}, nil
}

// BytesToInt8 convert bytes to int8, code 'b'
func BytesToInt8(b []byte, et EndianType) (n int8, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)

	switch len(b) {
	case 1:
		var x int8
		err = binary.Read(buf, endiantype, &x)
		n = x
	case 2:
		var x int16
		err = binary.Read(buf, endiantype, &x)
		n = int8(x)
	default:
		return 0, fmt.Errorf("input data length dont' match int8 need")
	}
	return
}

// Int82byte convert int8 to byte, code 'b'
// func Int82byte(n int8, et EndianType) (res []byte, err error) {
// 	return []byte{byte(n)}, nil
// }

// Int8ToBytes convert int8 to byte, code 'b', return 2 bytes
func Int8ToBytes(n int8, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	if err != nil {
		return nil, err
	}

	if buf.Len() < 2 {
		if et {
			res = append(res, []byte{0x00}...)
			res = append(res, buf.Bytes()...)
		} else {
			res = append(res, buf.Bytes()...)
			res = append(res, []byte{0x00}...)
		}
	}
	return res, nil
}

// BytesToUint8 convert bytes to uint8, code 'B'
func BytesToUint8(b []byte, et EndianType) (n uint8, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var x uint8
		err = binary.Read(buf, endiantype, &x)
		n = x
	case 2:
		var x uint16
		err = binary.Read(buf, endiantype, &x)
		n = uint8(x)
	default:
		return 0, fmt.Errorf("input data length dont' match uint8 need")
	}
	return
}

// Uint8ToBytes convert int8 to byte, code 'B', return 2 bytes
func Uint8ToBytes(n uint8, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	if err != nil {
		return nil, err
	}

	if buf.Len() < 2 {
		if et {
			res = append(res, []byte{0x00}...)
			res = append(res, buf.Bytes()...)
		} else {
			res = append(res, buf.Bytes()...)
			res = append(res, []byte{0x00}...)
		}
	}
	return res, nil
}

// BytesToInt16 convert bytes to int16, code 'h'
func BytesToInt16(b []byte, et EndianType) (n int16, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	if len(b) != 2 {
		return 0, fmt.Errorf("input data length dont' match int16 need")
	}
	err = binary.Read(buf, endiantype, &n)
	return n, err

}

// Int16ToBytes convert int16 to bytes, code 'h'
func Int16ToBytes(n int16, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToUint16 convert bytes to int16, code 'H'
func BytesToUint16(b []byte, et EndianType) (n uint16, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	if len(b) != 2 {
		return 0, fmt.Errorf("input data length dont' match int16 need")
	}
	err = binary.Read(buf, endiantype, &n)
	return n, err

}

// Uint16ToBytes convert uint16 to bytes, code 'H'
func Uint16ToBytes(n uint16, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToInt32 convert bytes to int32, code 'i/l'
func BytesToInt32(b []byte, et EndianType) (n int32, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	if len(b) != 4 {
		return 0, fmt.Errorf("input data length dont' match int32 need")
	}
	err = binary.Read(buf, endiantype, &n)
	return n, err
}

// Int32ToBytes convert int32 to bytes, code 'i/l'
func Int32ToBytes(n int32, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToUint32 convert bytes to int32, code 'I/L'
func BytesToUint32(b []byte, et EndianType) (n uint32, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	if len(b) != 4 {
		return 0, fmt.Errorf("input data length dont' match int32 need")
	}
	err = binary.Read(buf, endiantype, &n)
	return n, err
}

// Uint32ToBytes convert int32 to bytes, code 'I/L'
func Uint32ToBytes(n uint32, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToInt64 convert bytes to int64, code 'q'
func BytesToInt64(b []byte, et EndianType) (n int64, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	if len(b) != 8 {
		return 0, fmt.Errorf("input data length dont' match int64 need")
	}
	err = binary.Read(buf, endiantype, &n)
	return
}

// Int64ToBytes convert int16 to bytes, code 'q'
func Int64ToBytes(n int64, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToUint64 convert bytes to uint64, code 'Q'
func BytesToUint64(b []byte, et EndianType) (n uint64, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	if len(b) != 8 {
		return 0, fmt.Errorf("input data length dont' match int64 need")
	}
	err = binary.Read(buf, endiantype, &n)
	return
}

// Uint64ToBytes convert int64 to bytes, code 'Q'
func Uint64ToBytes(n uint64, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToFloat convert bytes to 32 bit float, code 'f'
func BytesToFloat(b []byte, et EndianType) (n float32, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)

	if len(b) != 4 {
		return 0, fmt.Errorf("input data length dont' match float need")
	}
	err = binary.Read(buf, endiantype, &n)
	return
}

// FloatToBytes convert 32 bit float to bytes, code 'f'
func FloatToBytes(n float32, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer([]byte{})
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

// BytesToDouble convert bytes to 32 bit float, code 'd'
func BytesToDouble(b []byte, et EndianType) (n float64, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)

	if len(b) != 8 {
		return 0, fmt.Errorf("input data length dont' match double need")
	}
	err = binary.Read(buf, endiantype, &n)
	return
}

// DoubleToBytes convert 64 bit float to bytes, code 'd'
func DoubleToBytes(n float64, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer([]byte{})
	err = binary.Write(buf, endiantype, n)
	return buf.Bytes(), err
}

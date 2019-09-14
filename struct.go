package gostruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
	gostruct

	Package gostruct is a formatter/conversion tool between values(can be readable) and binary byte(net or device register).
	it's a Python/Sturct like-style packager, provide almost same usage and and add new for another type.

	Format characters
		char	c type			Python type		Golang type		Standard size
		c		char				char		rune/string			1
		b		signed 				char/integer	int8			1
		B		unsigned char		integer			uint8			1
		?		_Bool				bool			boolt			1
		h		short				integer			int16			2
		H		unsigned short		integer			uint16			2
		i		int					integer			int				4
		I		unsigned int		integer			uint			4
		l		long				integer			int32			4
		L		unsigned long		long			uint32			4
		q		long long			long			int64			8
		Q		unsignedlonglong	long			uint64			8
		f		float				float			float32			4
		d		double				float			float64			8
		s		char[]				string
		p		char[]				string
		P		void *				long

*/

/*
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers
*/

const (
	//BigEndian Big endian //网络数据，Modbus, ARM,  MacOS,
	BigEndian EndianType = true
	//LittleEndian Little endian //X86 Windows, Linux，FreeBSD
	LittleEndian EndianType = false
	// 不同软件可能都不一致
)

//EndianType Data endian type, BigEndian or LittleEndian
type EndianType bool

type encoder struct {
	endian    EndianType
	Formatter string
}

func (e *encoder) calcsize() (n int) {
	dtysize := map[string]int{"c": 1, "b": 1, "B": 1, "?": 1, "h": 2, "H": 2, "i": 4, "I": 4, "l": 4, "L": 4, "q": 8, "Q": 8, "f": 4, "d": 8}
	for _, v := range e.Formatter {
		n += dtysize[string(v)]
	}
	return
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

func chars2byte(c string, size int) (res []byte, err error) {
	return
}

func bytes2char(b []byte) (c string, err error) {
	return
}

// input byte length must be 1
func byte2int8(b []byte, et EndianType) (n int8, err error) {
	if b == nil || len(b) == 0 {
		return 0, fmt.Errorf("input byte is empty or invaild")
	}
	if len(b) > 1 {
		b = b[0:1]
		err = fmt.Errorf("input byte beyong the int8 length")
	}
	var endiantype binary.ByteOrder
	endiantype = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := bytes.NewBuffer(b)
	err = binary.Read(buf, endiantype, &n)
	if err != nil {
		return 0, err
	}
	return
}

// input 1 or 2 []byte
func bytes2int8(b []byte, et EndianType) (n int8, err error) {
	var endiantype binary.ByteOrder
	endiantype = binary.LittleEndian
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
		return 0, fmt.Errorf("input byte beyong the int8 length")
	}
	return
}

//return 1 []byte
func int82byte(n int8, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder
	endiantype = binary.LittleEndian
	if et {
		endiantype = binary.BigEndian
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, endiantype, n)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//return 2 []byte
func int82bytes(n int8, et EndianType) (res []byte, err error) {
	var endiantype binary.ByteOrder
	endiantype = binary.LittleEndian
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

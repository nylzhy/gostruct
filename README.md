gostruct
==================


Package gostruct is a formatter/conversion tool between values(can be readable) and binary byte(net or device register). it's a Python/Sturct like-style packager, provide almost same usage and and add new for another type.

Format characters
-------------------------
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



data type range
-----------------------

- uint8       the set of all unsigned  8-bit integers (0 to 255)
- uint16      the set of all unsigned 16-bit integers (0 to 65535)
- uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
- uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)


- int8        the set of all signed  8-bit integers (-128 to 127)
- int16       the set of all signed 16-bit integers (-32768 to 32767)
- int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
- int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)


- float32     the set of all IEEE-754 32-bit floating-point numbers
- float64     the set of all IEEE-754 64-bit floating-point numbers

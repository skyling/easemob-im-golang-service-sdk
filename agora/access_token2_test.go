package agora

import (
	"fmt"
	"testing"
)

func TestPackInt(t *testing.T) {
	a := 55
	ret := PackInt(int16(a))
	fmt.Printf("%0x %d\n", ret, len(ret))
	aa, ret := UnPackUint16(ret)
	fmt.Printf("==a== %d %0x", aa, ret)
}

func TestPackString(t *testing.T) {
	strs := "123456"
	fmt.Println(strs[0:2], strs[2:])
	ret := PackString("habdsdfabdsdfabdsdfabdsdfabdsdfabdsdfabdsdfabdsdfabdsdf")
	ret = append(ret, PackUint32(123456)...)
	fmt.Printf("%0x\r\n", ret)
	str, ret := UnPackString(ret)
	i, ret := UnPackUint32(ret)
	fmt.Printf("%s %s %d", str, ret, i)
}

func TestZlibEncode(t *testing.T) {
	str := "abc123"
	by := ZlibEncode([]byte(str))
	de := ZlibDecode(by)
	fmt.Printf("%0x \n %s", by, de)
}

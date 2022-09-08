package agora

import (
	"bytes"
	"compress/zlib"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
	"time"
)

func PackUint16(n uint16) []byte {
	return PackInt(n)
}

func PackUint32(n uint32) []byte {
	return PackInt(n)
}

func UnPackUint16(data []byte) (uint16, []byte) {
	return binary.LittleEndian.Uint16(data[:2]), data[2:]
}

func UnPackUint32(data []byte) (uint32, []byte) {
	return binary.LittleEndian.Uint32(data[:4]), data[4:]
}

func PackInt(n interface{}) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, n)
	return bytesBuffer.Bytes()
}

func PackString(str string) []byte {
	ret := PackUint16(uint16(len(str)))
	ret = append(ret, []byte(str)...)
	return ret
}

func UnPackString(data []byte) (string, []byte) {
	lenS, data := UnPackUint16(data)
	return string(data[:lenS]), data[lenS:]
}

func PackMapUint32(data map[uint16]uint32) []byte {
	var keys []int
	var lens uint16
	for k := range data {
		keys = append(keys, int(k))
		lens += 1
	}
	ret := []byte{}
	sort.Ints(keys)
	for _, k := range keys {
		ret = append(ret, PackUint16(uint16(k))...)
		ret = append(ret, PackUint32(data[uint16(k)])...)
	}
	return append(PackUint16(lens), ret...)
}

func UnpackMapUint32(data []byte) (map[uint16]uint32, []byte) {
	lens, data := UnPackUint16(data)
	ret := map[uint16]uint32{}

	for i := 0; i < int(lens); i++ {
		var k uint16
		var v uint32
		k, data = UnPackUint16(data)
		v, data = UnPackUint32(data)
		ret[k] = v
	}
	return ret, data
}

func HmacSha256(key, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func ZlibEncode(src []byte) []byte {
	var bufs bytes.Buffer
	w, _ := zlib.NewWriterLevel(&bufs, -1)
	w.Write(src)
	defer w.Flush()
	w.Close()
	return bufs.Bytes()
}

func ZlibDecode(src []byte) []byte {
	buf := new(bytes.Buffer)
	buf.Write(src)
	read, _ := zlib.NewReader(buf)
	defer read.Close()
	deBuffer := new(bytes.Buffer)
	io.Copy(deBuffer, read)
	return deBuffer.Bytes()
}

func RandNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

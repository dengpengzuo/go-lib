package strutils

import (
    "math/rand"
    "time"
    "unsafe"
)

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const alphalen = len(alphanum)

func SubString(str string, start int) string {
    if start < len(str) {
        return str[start:]
    }
    return ""
}

func Bytes2String(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func String2Bytes(s string) []byte {
    x := (*[2]uintptr)(unsafe.Pointer(&s)) // stringstruct { str unsafe.Pointer, len int }
    h := [3]uintptr{x[0], x[1], x[1]}      // slice {addr unsafe.Pointer, len int , cap int }
    return *(*[]byte)(unsafe.Pointer(&h))
}

func RandomString(len int) string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    var bytes = make([]byte, len)
    n, e := r.Read(bytes)

    dir := false
    if n == len && e == nil {
        dir = true
    }
    if dir {
        for i, b := range bytes {
            bytes[i] = alphanum[b%byte(alphalen)]
        }
    } else {
        for i := 0; i < len; i++ {
            bytes[i] = alphanum[r.Int()%alphalen]
        }
    }

    return string(bytes)
}

package crand

//Charsets
const (
	CharSetDefault = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	CharSetAlpha   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	CharSetUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetLower   = "abcdefghijklmnopqrstuvwxyz"
)

//String returns a random string using CharSetDefault
func String(size int) string {
	return StringCharset(CharSetDefault, size)
}

//StringCharset generates a random string using the provided charset and size
func StringCharset(charSet string, size int) string {
	buf := make([]byte, size)
	charSetLen := uint(len(charSet))
	for i := 0; i < size; i++ {
		buf[i] = (charSet[Uint(charSetLen)])
	}
	return string(buf)
}

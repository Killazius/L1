package main

//var justString string
//
// вся строка v останется в памяти, даже если мы используем только первые 100 байт
// GC не сможет освободить память, потому что justString ссылается на часть v
// чтобы избежать этого, нужно скопировать нужную часть в новый слайс
// justString удерживает ненужную память от освобождения
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

var justString string

func createHugeString(n int) string {
	return string(make([]byte, n))
}
func someFunc() {
	v := createHugeString(1 << 10)
	buf := make([]byte, 100)
	copy(buf, v)
	justString = string(buf)
}

func main() {
	someFunc()
}

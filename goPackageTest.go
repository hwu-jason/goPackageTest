package goPackageTest
//https://golang.org/doc/code.html

//go build github.com/hwu-jason/goPackageTest
//go install github.com/hwu-jason/goPackageTest
//go test github.com/hwu-jason/goPackageTest
// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
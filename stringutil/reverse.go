// First statement must be package name. For every library package
// the package name should be the last name of the GOPATH
package stringutil

// return argument reversed from left to right
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i + 1, j - 1 {
        r[i], r[j]= r[j], r[i]
    }

    return string(r)
}

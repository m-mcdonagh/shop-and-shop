//go:build darwin
// +build darwin

package files

import "os"

var DataHome = "~/Library/shop-and-shop/"

func init() {
	if err := os.Mkdir(DataHome, 0755); err != nil && !os.IsExist(err) {
		panic(err)
	}
}

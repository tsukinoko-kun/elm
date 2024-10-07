//go:build windows && amd64

package elm

import (
	_ "embed"
)

var (
	//go:embed binary-for-windows-64-bit
	file []byte
)

const name = "elm.exe"

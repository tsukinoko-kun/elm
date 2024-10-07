//go:build linux && amd64

package elm

import (
	_ "embed"
)

var (
	//go:embed binary-for-linux-64-bit
	file []byte
)

const name = "elm"

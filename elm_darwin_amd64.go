//go:build darwin && amd64

package elm

import (
	_ "embed"
)

var (
	//go:embed binary-for-mac-64-bit
	file []byte
)

const name = "elm"

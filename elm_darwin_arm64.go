//go:build darwin && arm64

package elm

import (
	_ "embed"
)

var (
	//go:embed binary-for-mac-64-bit-ARM
	file []byte
)

const name = "elm"

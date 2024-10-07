# elm

Embedded Elm CLI to use in Go projects.

-   Use the `Init` function to initialize the elm binary. This will write the binary that is embedded in this package to a temporary file.
-   Use the `Dispose` function to remove the temporary file.
-   Use the `Command` function to create a `*exec.Cmd` that runs a command with the embedded elm binary.

```go
package main

import (
    "fmt"
    "os"

    "github.com/tsukinoko-kun/elm"
)

func main() {
    if err := elm.Init(); err != nil {
        panic(err)
    }
    defer elm.Dispose()

    cmd := elm.Command("make", "src/Main.elm", "--output=main.js")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println(err)
    }
}
```

# nixpacks-go

nixpacks-go is a golang sdk for [nixpacks](https://github.com/railwayapp/nixpacks).

## Dependencies

This sdk calls the shell script `nixpacks` so make sure to have it installed.
The sdk was build on top of version `1.4.2`

If you dont have nixpacks installed you can check the documentation at: https://nixpacks.com/docs/install

## Example

```go
package main

import (
    "context"
    "fmt"

    nixpacks "github.com/vano2903/nixpacks-go"
)

func main(){
    n, err := nixpacks.NewNixpacks()
    if err != nil {
        panic(err)
    }

    cmd, err := n.Build(context.Background(), nixpacks.BuildOptions{
        Path: "path/to/directory/to/build",
        Name: "image-name",
    })
    if err != nil {
        panic(err)
    }

    out, err := cmd.Result()
    if err != nil {
        panic(err)
    }

    fmt.Println(out.ImageName)
}
```

---

This sdk is still under development and is not ready for a production environment.

Most of the options to build and plan an image are implemented but there are still some missing (about cache mainly).

Feel free to contribute to this project and open issues if you find any bugs or have any suggestions.

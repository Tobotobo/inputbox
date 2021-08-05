# inputbox
golang inputbox

``` go
package main

import (
	"fmt"

	"github.com/Tobotobo/inputbox"
)

func main() {
	if result := inputbox.Show("prompt", "title", "default"); result != "" {
		fmt.Println(result)
	} else {
		fmt.Println("Cancel")
	}
}
```

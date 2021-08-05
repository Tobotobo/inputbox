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
![image](https://user-images.githubusercontent.com/46508541/128385661-fc5b1884-8680-4313-907d-ebadd5a72d5a.png)


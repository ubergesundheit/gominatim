# Gominatim - Interact with Nominatim in Go

My very own Nominatim wrapper

Very opinionated, use at your own risk!

## Usage example

```golang
package main

import (
	"fmt"

	"github.com/ubergesundheit/gominatim"
)

func main () {
 	geocoder, _ := gominatim.NewGominatim()

	res, err = geocoder.Search(gominatim.SearchParameters{Country: "Germany", City: "Hamburg"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
```

## License

[Apache-2.0](https://spdx.org/licenses/Apache-2.0.html)

package main

import (
	"flag"
	"fmt"
	"os"
)

/***
## Try equinox, use sample code.
1. Create app, and ID is : app_hdA9KZmCdYW
2. public key
```
-----BEGIN ECDSA PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEdOROI+hOlTvEdCAIek1Shrqw03dQBOPB
yGjR+yge9ymy46zZFgyS3yXoPi1ekrH4QxMSrk+0n38fPwCJK9FmGYaZ6X6ciEmX
PZm0cxkFGRSGbTUj9lwuQrHZy4ni3trS
-----END ECDSA PUBLIC KEY-----
```

***/

func init() {
	update := false
	print_line := false
	bili_vtuber := false
	flag.BoolVar(&print_line, "print", false, "Print a line...")
	flag.BoolVar(&update, "update", false, "Update app...")
	flag.BoolVar(&bili_vtuber, "vtuber", false, "get bili live room list....")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if update {
		equinoxUpdate() // update flag
		os.Exit(0)
	}

	if print_line {
		fmt.Println("Print a line.....")
		os.Exit(0)
	}

	if bili_vtuber {
		get_list()
		os.Exit(0)
	}
}

func main() {

}

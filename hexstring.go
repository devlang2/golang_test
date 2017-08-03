package main

import (
	"encoding/hex"
	"log"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	src := []byte("5f3d4526d15a37cf8243103b6004b3a13ff8abe735ecc788d4879f3bef34a92ce446cb97aed9350704351b27dfb7e851991ad101b0be39154165c61856be2f178513d057024eb8b628dfca77607742d68206c20667b6a54fb467bdbbd2df71ab1e4430bf4ad279db3d08332c55d12f05e1e996a46d11d9c753f845eb87b1c1189f0b3af3057c9dd657fbde1ac637cf62")
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(dst[:n])
}

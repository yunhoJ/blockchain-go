package main

import (
	"coin/explorer"
	"coin/rest"
)

func main() {
	go explorer.Start(4001)
	rest.Start(4000)

}

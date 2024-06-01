package cli

import (
	"coin/explorer"
	"coin/rest"
	"flag"
	"fmt"
	"os"
)

func useage() {
	fmt.Printf("welcome to 노마드 코드\n\n")
	fmt.Printf("use following command\n\n")
	fmt.Printf("explorer:	start html explorer\n")
	fmt.Printf("rest:		start rest api (recommand)\n\n")
	os.Exit(0)

}

func Start() {
	if len(os.Args) < 2 {
		useage()
	}
	// rest := flag.NewFlagSet("restr", flag.ExitOnError)
	// portFlag := rest.Int("port", 4000, "set the server port")
	// switch os.Args[1] {
	// case "explorer":
	// 	fmt.Println("start explorer")
	// case "rest":
	// 	rest.Parse(os.Args[2:])
	// 	fmt.Println("start rest api")

	// default:
	// 	useage()
	// }
	// if rest.Parsed() {
	// 	fmt.Print("test")
	// 	fmt.Println(*portFlag)
	// }

	port := flag.Int("port", 4000, "set the server port")
	port2 := flag.Int("port2", 4001, "set the server port")
	mode := flag.String("mode", "rest", "choose rest , html")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		if port == port2 {
			useage()
		}
		go rest.Start(*port)
		explorer.Start(*port2)
	default:
		useage()
	}
}

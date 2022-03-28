package main

// the main programme, includes the entrypoint.

import (
	"Tucker-package/exinit"
	pkg "Tucker-package/pkg"
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

/* GO MODULE!

1. project/yourPkg
2. go mod init project/yourPkg
3. mkdir pkg
4. code pkg/custompkg.go
5. mkdir programme
6. usepkg.go
7. go mod tidy
*/

func main() {
	pkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 2, 5, 6, 2, 8, 2, 4, 8, 1, 3, 5, 1, 0}
	graph := asciigraph.Plot(data)
	fmt.Print(graph, "\n")
	exinit.PrintD()
}

/*

f() d: 4					<----- Variables are initialised prior to the main function.
f() d: 5					<-----
exinit.init function 6		<-----
This is a custom package!
This is Github expkg Sample
 8.00 ┤    ╭╮ ╭╮
 7.00 ┤    ││ ││
 6.00 ┤  ╭╮││ ││
 5.00 ┤ ╭╯│││ ││ ╭╮
 4.00 ┤ │ │││╭╯│ ││
 3.00 ┼╮│ ││││ │╭╯│
 2.00 ┤╰╯ ╰╯╰╯ ││ │
 1.00 ┤        ╰╯ ╰╮
 0.00 ┤            ╰
f() d: 6					<----- Then finally runs PrintD().

*/

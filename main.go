package main

import (
	"github.com/Headcrab/unusedintf/pkg/unusedintf"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(unusedintf.Analyzer)
}

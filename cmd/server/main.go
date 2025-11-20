package main

import (
	"github.com/hogartr/go-hexagonal-template/di"
	"go.uber.org/fx"
)

func main() {
	fx.New(di.Module).Run()
}

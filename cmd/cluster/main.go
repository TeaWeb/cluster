package main

import (
	_ "github.com/TeaWeb/cluster/internal/bootstrap"
	"github.com/TeaWeb/cluster/internal/web"
)

func main() {
	web.Start()
}

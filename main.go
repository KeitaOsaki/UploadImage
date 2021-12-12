package main

import (
	"file-server/config"
	"file-server/pkg"
)

func main() {

	//サーバを起動
	addr := config.GetServerPort()
	pkg.Server.Run(addr)
}

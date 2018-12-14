package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xiaolehuigo/gameserver/config"
	"xiaolehuigo/gameserver/present"
)

func main() {
	config.Init()
	present.Listen()
}

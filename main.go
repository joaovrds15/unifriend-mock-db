package main

import (
	"unifriends-db-script/builder"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	builderIntance := builder.GetBuilder()
	director := builder.NewDirector(builderIntance)

	director.BuildUserResponse()

}

package main

import (
	"jobBoardApi/routers"
)

func main() {
	e := routers.JobPostingRoutes()
	e.Logger.Fatal(e.Start(":9090"))
}

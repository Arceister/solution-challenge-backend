package main

import (
	route "api-solution/Routes"
)

func main() {
	r := route.RouteHandler()
	r.Run()

}

//go:generate goagen bootstrap -d github.com/Ashtonian/goa-bug-repro/design

package main

import (
	"github.com/Ashtonian/goa-bug-repro/app"
	"github.com/goadesign/goa"
)

func main() {
	service := goa.New("cellar")
	c := NewBugController(service)
	app.MountBugController(service, c)
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

package main

import (
	"github.com/Ashtonian/goa-bug-repro/app"
	"github.com/goadesign/goa"
)

type BugController struct {
	*goa.Controller
}

func NewBugController(service *goa.Service) *BugController {
	return &BugController{Controller: service.NewController("BugController")}
}

func (c *BugController) Example(ctx *app.ExampleBugContext) error {
	return ctx.OK([]byte("Should return 400 missing param!"))
}

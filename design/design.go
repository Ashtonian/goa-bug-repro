package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("bug", func() {
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("bug", func() {
	BasePath("/api/bug")

	Action("example", func() {
		Routing(POST(""))
		Payload(Request)
		Response(OK)
	})
})

var Request = Type("exampleRequest", func() {
	Attribute("id", Integer, "")
	Attribute("startTime", DateTime, "")
	Attribute("endTime", DateTime, "")

	Required("id")
})

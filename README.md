# goa-bug-repro

This repository was made to reproduce a goa validation bug when only one field is required.

Assuming goa and its dependencies are install - To reproduce the bug run the api with `go build && ./go-bug-repro` and then `curl -d '{}' http://localhost:8080/api/bug`. The response should return something like `400 - Bad Request: missing required parameter: id` - because it's marked as a required field in the design code, but it returns a 200.

Upon further investigation - .Validate() call in unmarshall code inside of controllers.go is not generated if there is only one required primitive property and no actual validations for that property

This seems like it's a bug, unless it should just be explicitly documented somewhere. It's found at `goa/goagen/gen_app/writers.go:781`.

Following that up the stack we get to the comment in `goa/goagen/codegen/validation.go:269-274` which seems to allude to and attempt to deal with this edge case. From there, I had a bit of trouble following the code, so I'm just going to submit an issue for this one.
package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("feedpushr", func() {
	Title("The feedpushr API")
	Description("A simple feed aggregator daemon with sugar on top.")
	Contact(func() {
		Name("Nicolas Carlier")
		URL("https://github.com/ncarlier")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("feedpushr guide")
		URL("https://github.com/ncarlier/feedpusher/README.md")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/v1")

	Origin("http://swagger.goa.design", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

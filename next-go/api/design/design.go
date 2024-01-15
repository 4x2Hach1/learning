package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("api", func() {
	Title("Api Server Service")
	Description("Api server server for front.")
	Server("server", func() {
		Host("localhost", func() {
			URI("http://localhost:4000")
		})
	})
})

var _ = Service("server", func() {
	Description("Server Service for front.")

	Method("hello", func() {
		Payload(func() {
			Attribute("name", String, "Name")
			Required("name")
		})
		Result(String)
		HTTP(func() {
			GET("/hello/{name}")
			Response(StatusOK)
		})
	})
})

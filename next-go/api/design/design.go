package design

import (
	. "goa.design/goa/v3/dsl"
)

var JWTAuth = JWTSecurity("jwt", func() {
	Description("JWT based authentication")
	Scope("api:access", "API access")
})

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

	Method("users", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Required("token")
		})
		Result(ArrayOf(User))
		HTTP(func() {
			GET("/users")
			Response(StatusOK)
		})
	})
})

var User = ResultType("User", func() {
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "Name")
		Attribute("email", String, "Email")
	})
})

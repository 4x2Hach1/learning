package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
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
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("X-Shared-Secret")
		cors.Methods("GET", "POST", "PATCH", "DELETE")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})

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

	Method("login", func() {
		Payload(func() {
			Attribute("email", String, "Email")
			Attribute("password", String, "Password")
			Required("email", "password")
		})
		Result(LoginAuth)
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})

	// user ////////////////////////////////////////////////////////
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

	Method("userById", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Attribute("id", Int, "ID")
			Required("token", "id")
		})
		Result(User)
		HTTP(func() {
			GET("/user/{id}")
			Response(StatusOK)
		})
	})

	Method("newUser", func() {
		Payload(func() {
			Attribute("name", String, "Name")
			Attribute("email", String, "Email")
			Attribute("password", String, "Password")
			Required("name", "email", "password")
		})
		Result(Boolean)
		HTTP(func() {
			POST("/user")
			Response(StatusOK)
		})
	})

	Method("updateUser", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Attribute("name", String, "Name")
			Attribute("email", String, "Email")
			Required("token", "name", "email")
		})
		Result(Boolean)
		HTTP(func() {
			PATCH("/user")
			Response(StatusOK)
		})
	})

	// memory ////////////////////////////////////////////////////////
	Method("memories", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Required("token")
		})
		Result(ArrayOf(Memory))
		HTTP(func() {
			GET("/memories")
			Response(StatusOK)
		})
	})

	Method("memoryById", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Attribute("id", Int, "ID")
			Required("token", "id")
		})
		Result(Memory)
		HTTP(func() {
			GET("/memory/{id}")
			Response(StatusOK)
		})
	})

	Method("newMemory", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Attribute("title", String, "Title")
			Attribute("date", String, "Date", func() {
				Format(FormatDate)
			})
			Attribute("location", String, "Location")
			Attribute("detail", String, "Detail")
			Required("token", "title", "date", "location", "detail")
		})
		Result(Boolean)
		HTTP(func() {
			POST("/memory")
			Response(StatusOK)
		})
	})

	Method("deleteMemory", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Attribute("id", Int, "ID")
			Required("token", "id")
		})
		Result(Boolean)
		HTTP(func() {
			DELETE("/memory/{id}")
			Response(StatusOK)
		})
	})

	Method("updateMemory", func() {
		Security(JWTAuth, func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("token", String, "JWT token auth")
			Attribute("id", Int, "ID")
			Attribute("title", String, "Title")
			Attribute("date", String, "Date", func() {
				Format(FormatDate)
			})
			Attribute("location", String, "Location")
			Attribute("detail", String, "Detail")
			Required("token", "id", "title", "date", "location", "detail")
		})
		Result(Boolean)
		HTTP(func() {
			PATCH("/memory/{id}")
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

var LoginAuth = ResultType("LoginAuth", func() {
	Attributes(func() {
		Attribute("token", String, "Token")
		Attribute("user", User, "User")
	})
})

var Memory = ResultType("Memory", func() {
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("users_id", Int, "UsersId")
		Attribute("title", String, "Title")
		Attribute("date", String, "Date", func() {
			Format(FormatDate)
		})
		Attribute("location", String, "Location")
		Attribute("detail", String, "Detail")
	})
})

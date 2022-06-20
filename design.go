package design

import (
	. "goa.design/goa/v3/dsl"
)

// Main API declaration
var _ = API("clients", func() {
	Title("API for client")
	Description("This API manages client with CRUD operations")

	Server("Clients", func() {
		Host("localhost", func() {

			URL("http://loaclhost:8080/api/v1")
		})
	})
})

// Client Server declaration with two method and swagger API specification file

var _ = Service("CLient", func() {
	Description("The Client service allows access to client member")
	Method("add", func() {
		Payload(func() {
			Field(1, ClientID, string, "ClientID")
			Field(2, ClientName, string, "ClientName")
			Required("ClientID", "ClientName")
		})
		Result(Empty)
		Error("not_found", NotFound, "Client not found")
		HTTP(func() {
			POST("/api/v1/client/{ClientID}")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Payload(func() {
			Field(1, "ClientID", String, "Client ID")
			Required("ClientID")
		})
		Result(ClientManagement)
		Error("not_found", NotFound, "Client not found")
		HTTP(func() {
			GET("/api/v1/client")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Result(ColectionOf(ClientManagement))
		HTTP(func() {
			GET("/api/v1/client")
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")

})

// Client Management ia a custom ResultType used to configure views for our cutome type





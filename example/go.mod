module example

go 1.17

require (
	github.com/deepmap/oapi-codegen v1.10.1
	github.com/nautobot/go-nautobot v1.3.3
)

require github.com/google/uuid v1.3.0 // indirect

// replace github.com/nautobot/go-nautobot => ./../nautobot

.PHONY: gen-demo-proto
gen-demo-proto:
	@cwgo server -I ./idl --module mail --service demo_proto --idl ./idl/echo.proto

.PHONY: test-client
test-client:
	@go run ./cmd/client/client.go
.PHONY: test-server
test-server:
	@go run .

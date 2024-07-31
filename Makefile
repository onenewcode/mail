.PHONY: gen-demo-proto
gen-demo-proto:
	@cwgo server -I ./idl --module mail --service demo_proto --idl ./idl/echo.proto

.PHONY: test-client
test-client:
	@go run ./cmd/client/client.go
.PHONY: test-server
test-server:
	@go run .
.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift


.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m
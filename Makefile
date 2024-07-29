.PHONY: gen-demo-proto
gen-demo-proto:
	@cwgo server -I ./idl --module mail --service demo_proto --idl ./idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --server_name  frontend --module frontend --idl ../../idl/frontend/auth_page.proto

.PHONY: app-frontend
app-frontend: 
	@cd app/frontend && go run .

.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC  --server_name user --module  user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC  --server_name user --module  rpc_gen --I ../idl --idl ../idl/user.proto
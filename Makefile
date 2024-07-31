.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --server_name  frontend --module frontend --idl ../../idl/frontend/home.proto

.PHONY: app-frontend
app-frontend: 
	@cd app/frontend && go run .
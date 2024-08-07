export ROOT_MOD := ../..
.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --server_name  frontend --module frontend --idl ../../idl/frontend/auth_page.proto
.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC  --server_name user --module  user  --pass "-use   rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC  --server_name user --module  rpc_gen --I ../idl --idl ../idl/user.proto
.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --server_name product --module  rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --server_name product --module  product --pass "-use  rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --server_name cart --module  rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --server_name cart --module  cart --pass "-use  rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto

.PHONY: gen-checkout
gen-checkout: 
	@cd rpc_gen && cwgo client --type RPC --server_name checkout --module  rpc_gen  -I ../idl  --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --server_name checkout --module  checkout --pass "-use  rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto
.PHONY: gen-email
gen-email: 
	@cd rpc_gen && cwgo client --type RPC --server_name email --module  rpc_gen  -I ../idl  --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --server_name email --module  email --pass "-use  rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/email.proto
.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --server_name order --module  rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --server_name order --module  order --pass "-use  rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto
.PHONY: gen-payment
gen-payment: 
	@cd rpc_gen && cwgo client --type RPC --server_name payment --module  rpc_gen  -I ../idl  --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --server_name payment --module  payment --pass "-use  rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto


.PHONY: app-frontend
app-frontend: 
	@cd app/frontend && go run .

.PHONY: app-user
app-user: 
	@cd app/user && go run .
.PHONY: app-product
app-product: 
	@cd app/product && go run .
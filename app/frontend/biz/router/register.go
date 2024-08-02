

// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "frontend/biz/router/auth"
	cart "frontend/biz/router/cart"
	category "frontend/biz/router/category"
	checkout "frontend/biz/router/checkout"
	home "frontend/biz/router/home"
	order "frontend/biz/router/order"
	product "frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)
// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	checkout.Register(r)

	cart.Register(r)

	product.Register(r)

	category.Register(r)

	auth.Register(r)

	home.Register(r)
}
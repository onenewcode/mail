// Code generated by hertz generator.

package cart

import (
	"frontend/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Auth()}
}

func _addcartitemMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcartMw() []app.HandlerFunc {
	// your code...
	return nil
}

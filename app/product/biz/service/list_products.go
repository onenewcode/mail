

package service

import (
	"context"

	"product/biz/dal/mysql"
	"product/biz/model"
	product "rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
fuuct.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	c, err := model.GetProductsByCategoryName(mysql.DB, s.ctx, req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{Id: uint32(v.ID), Name: v.Name, Description: v.Description, Picture: v.Picture, Price: v.Price})
		}
	}

	return resp, nil
}

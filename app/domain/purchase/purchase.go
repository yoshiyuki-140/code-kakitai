package purchase

import (
	"context"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type PurchaseProduct struct {
	productID string
	count     int
}

func NewPurchaseProduct(productID string, count int) (*PurchaseProduct, error) {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return nil, errors.NewError("商品IDの値が不正です。")
	}

	// 購入数のバリデーション
	if count < 1 {
		return nil, errors.NewError("購入数の値が不正です。")
	}

	return &PurchaseProduct{
		productID: productID,
		count:     count,
	}, nil
}

func (p *PurchaseProduct) ProductID() string {
	return p.productID
}

func (p *PurchaseProduct) Count() int {
	return p.count
}

type PurchaseDomainService interface {
	PurchaseProducts(ctx context.Context, pps []PurchaseProduct) error
}
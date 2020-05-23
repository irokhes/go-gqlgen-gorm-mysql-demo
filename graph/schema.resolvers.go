package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-gqlgen-gorm-mysql-demo/graph/generated"
	"go-gqlgen-gorm-mysql-demo/graph/model"
	"log"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	order := model.Order{
		CustomerName: input.CustomerName,
		OrderAmount:  input.OrderAmount,
		Items:        mapItemsFromInput(input.Items),
	}
	log.Println(order)
	err := r.DB.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func mapItemsFromInput(itemsInput []*model.ItemInput) []model.Item {
	var items []model.Item
	for _, itemInput := range itemsInput {
		items = append(items, model.Item{
			ProductCode: itemInput.ProductCode,
			ProductName: itemInput.ProductName,
			Quantity:    itemInput.Quantity,
		})
	}
	return items
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	updateOrder := model.Order{
		ID:           orderID,
		CustomerName: input.CustomerName,
		Items:        mapItemsFromInput(input.Items),
	}
	r.DB.Save(&updateOrder)
	return &updateOrder, nil
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	r.DB.Where("order_id = ?", orderID).Delete(&model.Item{})
	r.DB.Where("order_id = ?", orderID).Delete(&model.Order{})

	return true, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	r.DB.Preload("Items").Find(&orders)

	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

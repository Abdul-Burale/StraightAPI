package models

import (
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	items := []Item{
		{
			Name: 		"Burger",
			Quantity:	1,
			Price:		5.99,
		},
		{
			Name:		"Pasta",
			Quantity:	2,
			Price:		12.99,
		},
	}

	order := Order{
		UserID:				"1",
		OrderNumber:		"0",
		Items:				items,
		TotalAmount:		18.98,
		Status:				"pending",
		CreatedAt:			time.Now(),
		UpdatedAt:			time.Now(),
	}

	if len(order.Items) != 2 {
		t.Errorf("TestCreateOrder: expected 2 items, got %d", len(items))
	}

	if order.UserID != "1" {
		t.Errorf("TestCreateOrder: expected UserID '1', got %v", order.UserID)
	}
}


func TestCreateItem(t *testing.T) {
	item := Item{
		Name: 		"Burger",
		Quantity: 	1,
		Price: 		5.99,
	}

	if item.Name != "Burger" {
		t.Errorf("TestItem: expected Name 'Burger', got %v", item.Name)
	}

	if item.Quantity != 1 {
		t.Errorf("TestItem: expected Quantity 1, got %d", item.Quantity)
	}

	if item.Price != 5.99 {
		t.Errorf("TestItem: expected Price 5.99, got %v", item.Price)
	}
}

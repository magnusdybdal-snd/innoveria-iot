package domain

import (
	"fmt"
	"time"
)

type OrderStatus string

const (
	OrderStatusNotInitialized OrderStatus = "not_initialized"
	OrderStatusRegistered     OrderStatus = "registered"
	OrderStatusPrinted        OrderStatus = "printed"
	OrderStatusStarted        OrderStatus = "started"
	OrderStatusFinished       OrderStatus = "finished"
	OrderStatusPostCalculated OrderStatus = "post_calculated"
	OrderStatusDelivered      OrderStatus = "delivered"
	OrderStatusHistorical     OrderStatus = "historical"
)

func ParseOrderStatus(value string) (OrderStatus, error) {
	status := OrderStatus(value)
	switch status {
	case OrderStatusNotInitialized,
		OrderStatusRegistered,
		OrderStatusPrinted,
		OrderStatusStarted,
		OrderStatusFinished,
		OrderStatusPostCalculated,
		OrderStatusDelivered,
		OrderStatusHistorical:
		return status, nil
	default:
		return "", fmt.Errorf("invalid order status: %q", value)
	}
}

type Order struct {
	ID                int
	OrderNumber       string
	PartID            string
	PartDescription   string // Human readable product name
	PlannedStartDate  time.Time
	PlannedFinishDate time.Time
	ActualStartDate   *time.Time
	ActualFinishDate  *time.Time
	Status            OrderStatus
	Priority          int
	ReceivedAt        time.Time
}

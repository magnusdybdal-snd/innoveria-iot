package domain

import (
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

type Order struct {
	ID                int64
	CompanyID         string
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

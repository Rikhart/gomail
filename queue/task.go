package queue

import (
	"context"
	"time"

	"github.com/mikestefanello/backlite"
)

// Tarea para enviar email de nuevo pedido
type NewOrderEmailTask struct {
	OrderID      string
	EmailAddress string
}

// Config() define cómo Backlite maneja la tarea
// Config debe devolver *backlite.QueueConfig
func (t NewOrderEmailTask) Config() backlite.QueueConfig {
	return backlite.QueueConfig{
		Name:        "NewOrderEmail",
		MaxAttempts: 5,
		Backoff:     5 * time.Second,
		Timeout:     10 * time.Second,
		Retention: &backlite.Retention{
			Duration:   6 * time.Hour,
			OnlyFailed: false,
			Data: &backlite.RetainData{
				OnlyFailed: true,
			},
		},
	}
}

// Función opcional para encolar la tarea
func EnqueueNewOrderEmail(bl *backlite.Client, orderID, email string) error {
	ctx := context.Background()
	task := NewOrderEmailTask{
		OrderID:      orderID,
		EmailAddress: email,
	}

	_, err := bl.Add(task).Ctx(ctx).At(time.Now()).Save()
	return err
}

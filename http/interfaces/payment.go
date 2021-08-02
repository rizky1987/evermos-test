package interfaces

import "evermos-test/database/entity"

type PaymentInterface interface {
	Create(e *entity.Payment) (bool, error)
	UpdateStatus(paymentCode, status string) error
}

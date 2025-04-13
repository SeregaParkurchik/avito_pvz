package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     string    `json:"Role"`
}

type PVZ struct {
	ID               uuid.UUID `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             string    `json:"city"`
}

type Receptions struct {
	ID       uuid.UUID `json:"id"`
	DateTime time.Time `json:"date_time"`
	PVZID    uuid.UUID `json:"pvzId"`
	Status   string    `json:"status"` // in_progress, close
}

type Product struct {
	ID           uuid.UUID `json:"id"` // айди самого товара
	DateTime     time.Time `json:"dateTime"`
	Type         string    `json:"type"`        // электроника, одежда, обувь
	ReceptionsID uuid.UUID `json:"receptionId"` //айди приемки
	PVZID        uuid.UUID `json:"pvzId"`       //айди пвз
}

type GetAllPVZRequest struct {
	StartDate time.Time `form:"startDate"`
	EndDate   time.Time `form:"endDate"`
	Page      int       `form:"page"`
	Limit     int       `form:"limit"`
}

type PVZWithReceptions struct {
	PVZ        PVZ                     `json:"pvz"`
	Receptions []ReceptionWithProducts `json:"receptions"`
}

type ReceptionWithProducts struct {
	Reception Receptions `json:"reception"`
	Products  []Product  `json:"products"`
}

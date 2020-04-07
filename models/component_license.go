package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// ComponentLicense is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type ComponentLicense struct {
	ID          uuid.UUID `json:"id" db:"id"` // Just keeping this because it seams that pop uses the ip column
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	LicenseID   string    `json:"-" db:"license_id"`
	ComponentID uuid.UUID `json:"-" db:"component_id"`
}

// ComponentLicenses is not required by pop and may be deleted
type ComponentLicenses []ComponentLicense

// String is not required by pop and may be deleted
func (c ComponentLicense) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// String is not required by pop and may be deleted
func (c ComponentLicenses) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *ComponentLicenses) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *ComponentLicenses) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *ComponentLicenses) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

package models

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// License is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type License struct {
	XMLName    xml.Name     `xml:"license" db:"-"`
	ID         uuid.UUID    `json:"id" db:"id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	SPDXID     string       `xml:"id" json:"spdx_id" db:"spdx_id"`
	Name       nulls.String `xml:"name" json:"name" db:"name"`
	URL        nulls.String `xml:"url" json:"url" db:"url"`
	Components Components   `json:"components,omitempty" many_to_many:"component_licenses"`
}

// Licenses is not required by pop and may be deleted
type Licenses []License

// String is not required by pop and may be deleted
func (c Licenses) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *License) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *License) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *License) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

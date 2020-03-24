package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Component is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Component struct {
	ID        uuid.UUID     `json:"id" db:"id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
	Type      string        `xml:"type,attr" json:"type" db:"type"`
	Publisher nulls.String  `xml:"publisher" json:"publisher" db:"publisher"`
	Group     nulls.String  `xml:"group" json:"group" db:"group"`
	Name      string        `xml:"name" json:"name" db:"name"`
	Version   string        `xml:"version" json:"version" db:"version"`
	Licenses  slices.String `xml:"licenses>licence>name" json:"licenses" db:"licenses"`
	Purl      string        `xml:"purl" json:"purl" db:"purl"`
}

// String is not required by pop and may be deleted
func (c Component) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Components is not required by pop and may be deleted
type Components []Component

// String is not required by pop and may be deleted
func (c Components) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Component) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Component) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Component) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

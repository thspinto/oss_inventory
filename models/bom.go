package models

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Bom is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Bom struct {
	XMLName    xml.Name   `xml:"bom" db:"-"`
	ID         uuid.UUID  `json:"id" db:"id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	Components Components `xml:"components>component" json:"components,omitempty" has_many:"components"`
	Project    string     `json:"project" db:"project"`
	Version    string     `json:"version" db:"version"`
}

// String is not required by pop and may be deleted
func (b Bom) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Boms is not required by pop and may be deleted
type Boms []Bom

// String is not required by pop and may be deleted
func (b Boms) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Bom) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Bom) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Bom) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

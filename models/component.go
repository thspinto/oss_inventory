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

// Component is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Component struct {
	XMLName   xml.Name     `xml:"component" db:"-"`
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Type      string       `xml:"type,attr" json:"type" db:"type"`
	Publisher nulls.String `xml:"publisher" json:"publisher" db:"publisher"`
	Group     nulls.String `xml:"group" json:"group" db:"group"`
	Name      string       `xml:"name" json:"name" db:"name"`
	Version   string       `xml:"version" json:"version" db:"version"`
	Licenses  Licenses     `xml:"licenses>license" json:"licenses" db:"licenses"`
	Purl      string       `xml:"purl" json:"purl" db:"purl"`
	BomID     int          `json:"-" db:"bom_id"`
	Bom       *Bom         `json:"bom,omitempty" belongs_to:"bom"`
}

// License is not required by pop and may be deleted
type License struct {
	XMLName xml.Name     `xml:"license"`
	ID      string       `xml:"id" json:"id"`
	Name    nulls.String `xml:"name" json:"name"`
	URL     nulls.String `xml:"url" json:"url"`
}

// String is not required by pop and may be deleted
func (c License) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Licenses is not required by pop and may be deleted
type Licenses []License

// String is not required by pop and may be deleted
func (c Licenses) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
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

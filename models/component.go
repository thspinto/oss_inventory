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
	XMLName   xml.Name     `xml:"component" json:"-" db:"-"`
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Type      string       `xml:"type,attr" json:"type" db:"type"`
	Publisher nulls.String `xml:"publisher" json:"publisher" db:"publisher"`
	Group     nulls.String `xml:"group" json:"group" db:"group"`
	Name      string       `xml:"name" json:"name" db:"name"`
	Version   string       `xml:"version" json:"version" db:"version"`
	Licenses  Licenses     `xml:"licenses>license" json:"licenses,omitempty" many_to_many:"component_licenses"`
	Purl      string       `xml:"purl" json:"purl" db:"purl"`
	BomID     uuid.UUID    `json:"-" db:"bom_id"`
	Bom       *Bom         `json:"bom,omitempty" belongs_to:"bom"`
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

// AssociateLicenses wraps the logic of creating the association between component and license
func (c *Component) AssociateLicenses(tx *pop.Connection, licenses Licenses) error {
	for _, license := range licenses {
		componentLicense := &ComponentLicense{
			LicenseID:   license.ID,
			ComponentID: c.ID,
		}
		err := tx.Create(componentLicense)
		if err != nil {
			return err
		}
	}
	return nil
}

// DissociateLicenses wraps the logic of creating the association between component and license
func (c *Component) DissociateLicenses(tx *pop.Connection, licenses Licenses) error {
	for _, license := range licenses {
		componentLicense := &ComponentLicense{}
		if err := tx.Where("component_id = ?", c.ID).Where("license_id = ?", license.ID).First(componentLicense); err != nil {
			return err
		}
		return tx.Destroy(componentLicense)
	}
	return nil
}

// TODO: Validate license on before validate

// AfterCreate add create and adds license associations
func (c *Component) AfterCreate(tx *pop.Connection) error {
	err := tx.Create(c.Licenses)
	if err != nil {
		return err
	}
	return c.AssociateLicenses(tx, c.Licenses)
}

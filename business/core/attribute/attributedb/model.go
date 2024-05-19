package attributedb

import (
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/attribute"
	"github.com/google/uuid"
)

// dbSubject represent the structure we need for moving data
// between the app and the database.
type dbAttribute struct {
	ID          uuid.UUID `db:"attribute_id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	Instance    int       `db:"instance"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBAttribute(att attribute.Attribute) dbAttribute {

	attri := dbAttribute{
		ID:          att.ID,
		Name:        att.Name,
		Type:        att.Type.Name(),
		Instance:    att.Instance,
		DateCreated: att.DateCreated.UTC(),
		DateUpdated: att.DateUpdated.UTC(),
	}

	return attri
}

func toCoreAttribute(dbAttribute dbAttribute) (attribute.Attribute, error) {

	typ, err := attribute.ParseType(dbAttribute.Type)
	if err != nil {
		return attribute.Attribute{}, fmt.Errorf("parse type: %w", err)
	}

	att := attribute.Attribute{
		ID:          dbAttribute.ID,
		Name:        dbAttribute.Name,
		Type:        typ,
		Instance:    dbAttribute.Instance,
		DateCreated: dbAttribute.DateCreated.In(time.Local),
		DateUpdated: dbAttribute.DateUpdated.In(time.Local),
	}

	return att, nil
}

func toCoreAttributeSlice(dbAttributes []dbAttribute) ([]attribute.Attribute, error) {
	atts := make([]attribute.Attribute, len(dbAttributes))
	for i, dbGa := range dbAttributes {
		var err error
		atts[i], err = toCoreAttribute(dbGa)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}
	return atts, nil
}

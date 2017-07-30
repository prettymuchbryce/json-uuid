package jsonuuid

import (
	"bytes"
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// UUID It's a RFC 4122 UUID
type UUID struct {
	uuid.NullUUID
}

// NewUUID instantiates a new, empty UUID
func NewUUID() UUID {
	return UUID{
		NullUUID: uuid.NullUUID{
			Valid: true,
			UUID:  uuid.NewV4(),
		}}
}

// FromString returns a UUID from a utf8 string
func FromString(s string) UUID {
	id, err := uuid.FromString(s)
	if err != nil {
		return UUID{}
	}
	return UUID{
		NullUUID: uuid.NullUUID{
			Valid: true,
			UUID:  id,
		},
	}
}

// FromBytes returns a UUID from bytes
func FromBytes(v []byte) UUID {
	id, err := uuid.FromBytes(v)
	if err != nil {
		return UUID{}
	}
	return UUID{
		NullUUID: uuid.NullUUID{
			Valid: true,
			UUID:  id,
		},
	}
}

// MarshalJSON marshalls the NullUUID as nil or the nested UUID
func (u UUID) MarshalJSON() ([]byte, error) {
	if !u.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(u.UUID)
}

// UnmarshalJSON unmarshalls a NullUUID
func (u *UUID) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		u.NullUUID = uuid.NullUUID{
			Valid: false,
		}
		return nil
	}

	if err := json.Unmarshal(b, &u.UUID); err != nil {
		return err
	}
	u.Valid = true

	return nil
}

package core

import (
	"time"

	"github.com/google/uuid"
)

type ID int64
type PublicID uuid.UUID
type OwnershipID ID
type IdentityDocument string
type Email string
type Enabled bool
type CreatedAt time.Time
type UpdatedAt time.Time
type DeletedAt *time.Time
type Version int64

func (c CreatedAt) String() string {
	return time.Time(c).Format(time.RFC3339Nano)
}

func (u UpdatedAt) String() string {
	return time.Time(u).Format(time.RFC3339Nano)
}

func (c CreatedAt) Format() string {
	return time.Time(c).Format(time.RFC3339)
}

func (u UpdatedAt) Format() string {
	return time.Time(u).Format(time.RFC3339)
}

func (v Version) Int() int64 {
	return int64(v)
}

func (i ID) Int() int64 {
	return int64(i)
}

func (p PublicID) String() string {
	return uuid.UUID(p).String()
}

func (o OwnershipID) Int() int64 {
	return int64(o)
}

func (i IdentityDocument) String() string {
	return string(i)
}

func (e Email) String() string {
	return string(e)
}

func (e Enabled) Bool() bool {
	return bool(e)
}

func NewID() ID {
	return ID(0)
}

func NewOwnershipID() OwnershipID {
	return OwnershipID(0)
}

func NewEnabled() Enabled {
	return Enabled(true)
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func NewVersion() Version {
	return Version(1)
}

func NewPublicID() PublicID {
	return PublicID(uuid.New())
}

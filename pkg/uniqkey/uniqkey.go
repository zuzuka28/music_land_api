package uniqkey

import (
	"strconv"
	"sync/atomic"
)

var lastID int32 = 1

// DataKey used to maintain the uniqueness of properties between different packages.
type DataKey struct{ id int32 }

func (d DataKey) String() string {
	return strconv.Itoa(int(d.id))
}

// Gen generates new unique key for property.
func Gen() DataKey { return DataKey{id: atomic.AddInt32(&lastID, 1)} }

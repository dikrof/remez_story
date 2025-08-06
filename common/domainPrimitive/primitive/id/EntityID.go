package id

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// EntityID — общий доменный примитив идентификатора (int64 > 0).
type EntityID struct{ v int64 }

// NewEntityID создаёт валидный ID (>0).
func NewEntityID(v int64) (EntityID, error) {
	if v <= 0 {
		return EntityID{}, errors.New("EntityID must be > 0")
	}
	return EntityID{v: v}, nil
}

// EntityIDFrom парсит ID из строки.
func EntityIDFrom(s string) (EntityID, error) {
	s = strings.TrimSpace(s)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return EntityID{}, err
	}
	return NewEntityID(n)
}

func (x EntityID) Int64() int64 { return x.v }

func (x EntityID) IsZero() bool { return x.v == 0 }

func (x EntityID) String() string { return fmt.Sprintf("%d", x.v) }

func (x EntityID) MarshalJSON() ([]byte, error) { return json.Marshal(x.v) }

func (x *EntityID) UnmarshalJSON(b []byte) error {
	var n int64
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	v, err := NewEntityID(n)
	if err != nil {
		return err
	}
	*x = v
	return nil
}

func (x EntityID) Value() (driver.Value, error) { return x.v, nil }

func (x *EntityID) Scan(src any) error {
	switch t := src.(type) {
	case int64:
		v, err := NewEntityID(t)
		if err != nil {
			return err
		}
		*x = v
		return nil
	case []byte:
		n, err := strconv.ParseInt(string(t), 10, 64)
		if err != nil {
			return err
		}
		v, err := NewEntityID(n)
		if err != nil {
			return err
		}
		*x = v
		return nil
	case string:
		n, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return err
		}
		v, err := NewEntityID(n)
		if err != nil {
			return err
		}
		*x = v
		return nil
	default:
		return fmt.Errorf("EntityID: unsupported Scan type %T", src)
	}
}

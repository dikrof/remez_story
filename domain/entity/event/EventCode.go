package event

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var codeRe = regexp.MustCompile(`^[A-Z0-9_]{1,64}$`)

// EventCode — код события (уникален в справочнике).
type EventCode struct{ v string }

func NewCode(s string) (EventCode, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ToUpper(s)
	if !codeRe.MatchString(s) {
		return EventCode{}, errors.New("EventCode must match ^[A-Z0-9_]{1,64}$")
	}
	return EventCode{v: s}, nil
}

func MustCode(s string) EventCode {
	c, err := NewCode(s)
	if err != nil {
		panic(err)
	}
	return c
}

func (c EventCode) String() string               { return c.v }
func (c EventCode) IsZero() bool                 { return c.v == "" }
func (c EventCode) MarshalJSON() ([]byte, error) { return json.Marshal(c.v) }
func (c *EventCode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewCode(s)
	if err != nil {
		return err
	}
	*c = x
	return nil
}
func (c EventCode) Value() (driver.Value, error) { return c.v, nil }
func (c *EventCode) Scan(src any) error {
	switch t := src.(type) {
	case string:
		x, err := NewCode(t)
		if err != nil {
			return err
		}
		*c = x
		return nil
	case []byte:
		x, err := NewCode(string(t))
		if err != nil {
			return err
		}
		*c = x
		return nil
	default:
		return fmt.Errorf("EventCode: unsupported Scan type %T", src)
	}
}

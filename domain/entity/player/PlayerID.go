package player

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// PlayerID — доменный примитив идентификатора игрока.
// Можно хранить UUID или любой валидированный ключ до 128 символов.
type PlayerID struct{ v string }

// New создаёт ID с валидацией (не пустой, ограничение длины).
func NewPlayerID(s string) (PlayerID, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return PlayerID{}, errors.New("PlayerID must be non-empty")
	}
	if len(s) > 128 {
		return PlayerID{}, errors.New("PlayerID too long")
	}
	return PlayerID{v: s}, nil
}

// Must — как New, но паникует при ошибке.
func MustPlayerID(s string) PlayerID {
	id, err := NewPlayerID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// String возвращает строковое значение.
func (id PlayerID) String() string { return id.v }

// IsZero сообщает, что ID пуст.
func (id PlayerID) IsZero() bool { return id.v == "" }

func (id PlayerID) MarshalJSON() ([]byte, error) { return json.Marshal(id.v) }

func (id *PlayerID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewPlayerID(s)
	if err != nil {
		return err
	}
	*id = x
	return nil
}

func (id PlayerID) Value() (driver.Value, error) { return id.v, nil }

func (id *PlayerID) Scan(src any) error {
	switch t := src.(type) {
	case string:
		x, err := NewPlayerID(t)
		if err != nil {
			return err
		}
		*id = x
		return nil
	case []byte:
		x, err := NewPlayerID(string(t))
		if err != nil {
			return err
		}
		*id = x
		return nil
	default:
		return fmt.Errorf("PlayerID: unsupported Scan type %T", src)
	}
}

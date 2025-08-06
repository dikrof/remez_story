package node

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SceneLabel — "мягкая" метка сцены.
// Может быть пустой. Ограничение длины — 128 символов.
type SceneLabel struct{ v string }

// NewSceneLabel создаёт SceneLabel с валидацией длины.
func NewSceneLabel(s string) (SceneLabel, error) {
	s = strings.TrimSpace(s)
	if len(s) > 128 {
		return SceneLabel{}, errors.New("SceneLabel too long")
	}
	return SceneLabel{v: s}, nil
}

// String возвращает строковое значение метки.
func (l SceneLabel) String() string { return l.v }

// IsZero сообщает, что метка пустая.
func (l SceneLabel) IsZero() bool { return l.v == "" }

func (l SceneLabel) MarshalJSON() ([]byte, error) { return json.Marshal(l.v) }

func (l *SceneLabel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewSceneLabel(s)
	if err != nil {
		return err
	}
	*l = x
	return nil
}

func (l SceneLabel) Value() (driver.Value, error) { return l.v, nil }

func (l *SceneLabel) Scan(src any) error {
	switch t := src.(type) {
	case string:
		x, err := NewSceneLabel(t)
		if err != nil {
			return err
		}
		*l = x
		return nil
	case []byte:
		x, err := NewSceneLabel(string(t))
		if err != nil {
			return err
		}
		*l = x
		return nil
	default:
		return fmt.Errorf("SceneLabel: unsupported Scan type %T", src)
	}
}

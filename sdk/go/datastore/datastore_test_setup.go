/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package datastore

import (
	"encoding/json"
	"fmt"
	"time"
)

type Friend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NamedString string

const (
	NamedStringOne   = "one"
	NamedStringTwo   = "two"
	NamedStringThree = "three"
)

func (p *NamedString) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*p = NamedString(v)
		return nil
	case []byte:
		*p = NamedString(string(v))
		return nil
	default:
		return fmt.Errorf("cannot scan NamedString from non-string value: %T", value)
	}
}

type NamedStrings []NamedString

func (p *NamedStrings) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		var items []string
		err := json.Unmarshal([]byte(v), &items) // if stored as JSON array
		if err != nil {
			return err
		}
		*p = make([]NamedString, len(items))
		for i, s := range items {
			(*p)[i] = NamedString(s)
		}
		return nil
	case []byte:
		return p.Scan(string(v))
	default:
		return fmt.Errorf("cannot scan %T into NamedStrings", src)
	}
}

type TestObject struct {
	Name string `json:"name"`
	// Omit empty is required to test missing values
	Value             int                     `json:"value,omitempty"`
	Age               int                     `json:"age"`
	NickNames         []string                `json:"nickNames"`
	Friend            Friend                  `json:"friend"`
	Friends           []Friend                `json:"friends"`
	FriendPointers    []*Friend               `json:"friendPointers"`
	Amap              map[string]interface{}  `json:"amap"`
	AmapPointer       *map[string]interface{} `json:"amapPointer"`
	AmapString        map[string]string       `json:"amapString"`
	AmapStringPointer *map[string]string      `json:"amapStringPointer"`
	FriendPointer     *Friend                 `json:"friendPointer"`
	CreatedAt         time.Time               `json:"createdAt"`
	NamedType         NamedString             `json:"namedType"`
	NamedTypes        NamedStrings            `json:"namedTypes"`
}

func (t TestObject) Indexes() []Index {
	return []Index{
		{
			Fields:    []string{"value"},
			Ascending: false,
		},
	}
}

func (t TestObject) GetId() string {
	return t.Name
}

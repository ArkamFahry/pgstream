package models

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"time"
)

const (
	EventOperationSnapshot = "snapshot"
	EventOperationInsert   = "insert"
	EventOperationUpdate   = "update"
	EventOperationDelete   = "delete"
	EventOperationTruncate = "truncate"
	EventOperationUnknown  = "unknown"
)

type Event struct {
	Id        string `json:"id"`
	Database  string `json:"database"`
	Schema    string `json:"schema"`
	Table     string `json:"table"`
	Operation string `json:"operation"`
	Data      struct {
		Before map[string]any `json:"before"`
		After  map[string]any `json:"after"`
		Diff   map[string]any `json:"diff"`
	} `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *Event) SetDefaults() {
	e.Id = ulid.Make().String()
}

func (e *Event) Subject() string {
	return fmt.Sprintf("%s.%s.%s.%s", e.Database, e.Schema, e.Table, e.Operation)
}

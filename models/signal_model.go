package models

import "time"

const (
	SignalTable = "pgwarp.signals"
)

const (
	SignalTypeSnapshotStart  = "snapshot.start"
	SignalTypeSnapshotPause  = "snapshot.pause"
	SignalTypeSnapshotResume = "snapshot.resume"
	SignalTypeSnapshotEnd    = "snapshot.end"
)

type Signal struct {
	Id        string    `json:"id" db:"id"`
	Type      string    `json:"type" db:"type"`
	Database  string    `json:"database" db:"database"`
	Schema    string    `json:"schema" db:"schema"`
	Tables    []string  `json:"tables" db:"tables"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type SnapshotPause struct {
	Id        string `json:"id" db:"id"`
	SignalId  string `json:"signal_id" db:"signal_id"`
	Point     string `json:"pause_point" db:"point"`
	ResumedAt string `json:"resumed_at" db:"resumed_at"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

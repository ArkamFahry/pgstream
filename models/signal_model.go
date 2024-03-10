package models

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
	Id   string         `json:"id" db:"id"`
	Type string         `json:"type" db:"type"`
	Data map[string]any `json:"data" db:"data"`
}

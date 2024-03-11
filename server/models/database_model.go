package models

const (
	SnapshotModeNever       = "never"
	SnapshotModeInitial     = "initial"
	SnapshotModeInitialOnly = "initial_only"
	SnapshotModeCustom      = "custom"

	SnapshotTypeNormal      = "normal"
	SnapshotTypeIncremental = "incremental"

	EventAll      = "*"
	EventSnapshot = "snapshot"
	EventInsert   = "insert"
	EventUpdate   = "update"
	EventDelete   = "delete"
	EventTruncate = "truncate"
	EventUnknown  = "unknown"

	StreamTypePersistent = "persistent"
	StreamTypeTransient  = "transient"
)

type Database struct {
	Name                    string   `json:"name" mapstructure:"name"`
	PostgresUrl             string   `json:"postgres_url" mapstructure:"postgres_url"`
	PostgresReplicationSlot string   `json:"postgres_replication_slot" mapstructure:"postgres_replication_slot"`
	NatsUrls                []string `json:"nats_urls" mapstructure:"nats_urls"`
	Schema                  struct {
		Includes []struct {
			Name  string `json:"name" mapstructure:"name"`
			Table struct {
				Includes []struct {
					Name       string   `json:"name" mapstructure:"name"`
					Columns    []string `json:"columns" mapstructure:"columns"`
					PrimaryKey []string `json:"primary_key" mapstructure:"primary_key"`
					Events     []string `json:"events" mapstructure:"events"`
					Snapshot   bool     `json:"snapshot" mapstructure:"snapshot"`
					StreamType string   `json:"stream_type" mapstructure:"stream_type"`
				} `json:"includes" mapstructure:"includes"`
				Excludes []string `json:"excludes" mapstructure:"excludes"`
			} `json:"table" mapstructure:"table"`
		} `json:"includes" mapstructure:"includes"`
		Excludes []string `json:"excludes" mapstructure:"excludes"`
	} `json:"schema" mapstructure:"schema"`
	SnapshotMode string `json:"snapshot_mode" mapstructure:"snapshot_mode"`
	SnapshotType string `json:"snapshot_type" mapstructure:"snapshot_type"`
}

type DatabaseStructure struct {
	Name    string `json:"name"`
	Schemas []struct {
		Name   string `json:"name"`
		Tables []struct {
			Name       string   `json:"name"`
			PrimaryKey []string `json:"primary_key"`
			Columns    []string `json:"columns"`
		} `json:"tables"`
	} `json:"schemas"`
}

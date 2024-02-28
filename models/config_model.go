package models

const (
	SnapshotModeNever       = "never"
	SnapshotModeInitial     = "initial"
	SnapshotModeInitialOnly = "initial_only"
	SnapshotModeCustom      = "custom"

	SnapshotTypeNormal      = "normal"
	SnapshotTypeIncremental = "incremental"

	StreamTypePersistent = "persistent"
	StreamTypeTransient  = "transient"
)

type DataConnection struct {
	Databases []Database `json:"databases" mapstructure:"databases"`
}

type Database struct {
	PostgresUrl   string         `json:"postgres_url" mapstructure:"postgres_url"`
	NatsUrls      []string       `json:"nats_urls" mapstructure:"nats_urls"`
	Database      string         `json:"database" mapstructure:"database"`
	Schemas       []string       `json:"schemas" mapstructure:"schemas"`
	TableIncludes []TableInclude `json:"table_includes" mapstructure:"table_includes"`
	TableExcludes []TableExclude `json:"table_excludes" mapstructure:"table_excludes"`
	SnapshotMode  string         `json:"snapshot_mode" mapstructure:"snapshot_mode"`
	SnapshotType  string         `json:"snapshot_type" mapstructure:"snapshot_type"`
}

type TableInclude struct {
	Schema      string   `json:"schema" mapstructure:"schema"`
	Name        string   `json:"name" mapstructure:"name"`
	PrimaryKeys []string `json:"primary_keys" mapstructure:"primary_keys"`
	Columns     []string `json:"columns" mapstructure:"columns"`
	Snapshot    bool     `json:"snapshot" mapstructure:"snapshot"`
	StreamType  string   `json:"stream_type" mapstructure:"stream_type"`
}

type TableExclude struct {
	Schema string `json:"schema" mapstructure:"schema"`
	Name   string `json:"name" mapstructure:"name"`
}

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

type Config struct {
	Databases []struct {
		Name        string   `json:"name" mapstructure:"name"`
		PostgresUrl string   `json:"postgres_url" mapstructure:"postgres_url"`
		NatsUrls    []string `json:"nats_urls" mapstructure:"nats_urls"`
		Schema      struct {
			Includes []struct {
				Name  string `json:"name" mapstructure:"name"`
				Table struct {
					Includes []struct {
						Name        string   `json:"name" mapstructure:"name"`
						PrimaryKeys []string `json:"primary_keys" mapstructure:"primary_keys"`
						Columns     []string `json:"columns" mapstructure:"columns"`
						Events      []string `json:"events" mapstructure:"events"`
						Snapshot    bool     `json:"snapshot" mapstructure:"snapshot"`
						StreamType  string   `json:"stream_type" mapstructure:"stream_type"`
					} `json:"includes" mapstructure:"includes"`
					Excludes []struct {
						Id   string `json:"id" mapstructure:"id"`
						Name string `json:"name" mapstructure:"name"`
					} `json:"excludes" mapstructure:"excludes"`
				} `json:"table" mapstructure:"table"`
			} `json:"includes" mapstructure:"includes"`
			Excludes []string `json:"excludes" mapstructure:"excludes"`
		} `json:"schema" mapstructure:"schema"`
		SchemaExcludes []string `json:"schema_excludes" mapstructure:"schema_excludes"`
		SnapshotMode   string   `json:"snapshot_mode" mapstructure:"snapshot_mode"`
		SnapshotType   string   `json:"snapshot_type" mapstructure:"snapshot_type"`
	} `json:"databases" mapstructure:"databases"`
}

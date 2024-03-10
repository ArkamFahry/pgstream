# PGWarp

**PGWarp** is a service which captures [Postgres](https://www.postgresql.org/) SnapShot, Incremental-SnapShot and [CDC](https://en.wikipedia.org/wiki/Change_data_capture) Events and publishes them to [NATS](https://nats.io/). 
**PGWarp** turns Postgres into a real-time reactive event stream.

## Features

- CDC based capture of all database changes in Postgres and publishing them to NATS streams or NATS pub/sub.
- SnapShot or Incremental-SnapShot of historical data from Postgres and publishing them to a NATS stream.
- Non locking Incremental-SnapShot without locking tables.
- Incremental-SnapShot of table while capturing table CDC events both run in parallel at the same time and returns data as a single stream of events.

## References

- [DBLog](https://netflixtechblog.com/dblog-a-generic-change-data-capture-framework-69351fb9099b)
- [pgcapture](https://github.com/replicase/pgcapture)
- [Debezium](https://debezium.io/)
- [WAL-Listener](https://github.com/ihippik/wal-listener)
- [DBConvert Streams](https://stream.dbconvert.com/)
# PGWarp

**PGWarp** is a service which captures Postgres SnapShot, Incremental-SnapShot and CDC Events and publishes them to NATS. 
**PGWarp** turns Postgres into a near real-time reactive event stream.

## Features

- CDC based capture of all database changes in Postgres and publishing them to NATS streams or NATS pub/sub.
- SnapShot or Incremental-SnapShot of historical data from Postgres and publishing them to a NATS stream.
- Non locking Incremental-SnapShot without locking tables.
- Incremental-SnapShot of table while capturing table CDC events both run in parallel at the same time and returns data as a single stream of events.
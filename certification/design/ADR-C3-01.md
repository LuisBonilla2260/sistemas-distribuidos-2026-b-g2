# ADR-C3-01: Two availability zones with managed, replicated data stores

## Status

Proposed

## Context

Section 2 of the design note (Reliability row) shows that every stage runs one PostgreSQL container and one MongoDB container on a single machine, and that the MongoDB replica-set size and the hosting are still pending decision P-01. Section 4 sets an RTO of 15 minutes and an RPO of 5 minutes for the Clinical MongoDB when an availability zone fails. With one copy of each store in one place, a zone failure would lose the clinical records written since the last archive and stop every service until someone restores by hand, which breaks both targets.

## Decision

Deploy the platform in the AP-Singapore region across two availability zones. PostgreSQL moves to RDS for PostgreSQL in primary/standby mode with the standby in the second zone and synchronous replication. Clinical MongoDB moves to DDS as a three-member replica set spread over both zones. Redis uses DCS master/standby and RabbitMQ uses DMS in cluster mode. The CCE worker nodes are split evenly between the two zones, and each service keeps at least one pod per zone. CBR takes scheduled backups of RDS and DDS and copies them to OBS for the full-region case.

## Alternatives considered

- Self-managed PostgreSQL and MongoDB containers on ECS in two zones: rejected because the team would have to build and test failover, backups and patching itself, which is exactly the part still pending in `13-operations/backup-and-recovery.md`.
- Single zone with frequent backups only: rejected because a restore from an archive takes far longer than 15 minutes and loses everything since the last backup, so it cannot meet the 5-minute RPO.

## Consequences

- Positive: a zone failure becomes an automatic failover instead of a manual restore, and the team stops maintaining replication scripts.
- Negative: the standby and the extra replica roughly double the database cost, and the shared PostgreSQL instance remains a single blast radius for four owners.

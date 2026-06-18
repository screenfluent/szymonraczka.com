# Database

Use this when changing SQL, SQLite, migrations, schema, indexes, constraints, backup/restore, or import/export.

## SQL generation

If the project uses generators:

```sh
sqlc generate
```

Then:

```sh
go test ./...
```

## SQLite checks

```sh
sqlite3 app.db 'PRAGMA integrity_check;'
sqlite3 app.db 'PRAGMA quick_check;'
sqlite3 app.db 'PRAGMA foreign_key_check;'
```

## Migrations

Test:

- fresh database from zero
- upgrade from previous schema
- data preserved after migration
- rollback only if project supports rollback

Use project migration tool. Do not edit schema manually outside the migration path.

## Query performance

```sh
sqlite3 app.db 'EXPLAIN QUERY PLAN SELECT ...;'
```

## Backup/restore

A backup is not proven until restore is tested.

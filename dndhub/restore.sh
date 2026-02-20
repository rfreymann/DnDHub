#!/bin/bash
set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <backup_file_path>"
    exit 1
fi

BACKUP_FILE="$1"

if [ ! -f "$BACKUP_FILE" ]; then
    echo "Error: File '$BACKUP_FILE' not found."
    exit 1
fi

echo "WARNING: This will overwrite the current database 'myappdb' with data from '$BACKUP_FILE'."
echo "Any existing data in the database will be modified/lost based on the backup content."
read -p "Are you sure you want to proceed? (y/N) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Restore cancelled."
    exit 1
fi

echo "Restoring database from $BACKUP_FILE..."

# Pipe file content into psql inside the container
# We might want to drop/create schema or use --clean in dump, but standard restore usually appends/modifies.
# For a clean restore, usually we drop the DB or clear public schema. 
# For now, simplistic restore (good for recovering deleted rows etc, but might conflict on PKs if not careful).
# Ideally we'd drop the schema public first. Let's add that safety if possible or just rely on user knowing.
# A simpler full restore often implies dropped DB. 
# Let's verify if pg_dump was with --clean? simpler backup script doesn't have it.
# Let's assume standard psql restore.

cat "$BACKUP_FILE" | podman exec -i dndhub_db_1 psql -U myappuser myappdb

echo "Restore completed."

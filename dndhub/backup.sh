#!/bin/bash
set -e

# Ensure backups directory exists
mkdir -p backups

# Generate timestamped filename
FILENAME="backups/dndhub_$(date +%Y-%m-%d_%H%M%S).sql"

echo "Starting backup to $FILENAME..."

# Execute pg_dump inside the container and redirect output to host file
# Note: Using -U myappuser myappdb as defined in podman-compose.yml
podman exec dndhub_db_1 pg_dump -U myappuser myappdb > "$FILENAME"

echo "Backup completed successfully: $FILENAME"
ls -lh "$FILENAME"

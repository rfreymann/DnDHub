# Database Migrations for DnDHub

This folder contains SQL migration files for the DnDHub backend.  
Each migration file is versioned and applied in order.

## Structure

migrations/
├── 001_init.sql   # initial schema (users, franchises, unique_workers)
├── 002_*.sql      # future changes

We use sequential numbering (`001_`, `002_`, …) to keep migrations ordered.

---

## Option A: Run manually with psql

1. Create your database (if not already):

    createdb dndhub

2. Apply migrations:

    psql -d dndhub -f migrations/001_init.sql

3. Verify tables exist:

    psql -d dndhub -c "\dt"

---

## Option B: Run automatically with golang-migrate

We recommend [golang-migrate](https://github.com/golang-migrate/migrate) for production.

1. Install migrate:

    brew install golang-migrate   # macOS  
    apt install migrate           # Debian/Ubuntu  

    Or download from releases: https://github.com/golang-migrate/migrate/releases

2. Run migrations:

    migrate -path migrations -database "$DATABASE_URL" up

   Example `DATABASE_URL` for Postgres:

    postgres://user:password@localhost:5432/dndhub?sslmode=disable

3. Roll back last migration:

    migrate -path migrations -database "$DATABASE_URL" down 1

---

## Notes

- `001_init.sql` creates:
  - `users`
  - `franchises`
  - `unique_workers`

- Always create a **new numbered file** for schema changes (never edit old migrations).  
- Keep this folder in Git for team sync.

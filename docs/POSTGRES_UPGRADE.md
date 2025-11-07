# PostgreSQL Docker 18+ — Volume mount & upgrade notes

This note explains an important change when using the official PostgreSQL Docker images 18+ and the safe way to migrate existing data.

Summary
- Recent postgres Docker images (18+) place server clusters in versioned directories under `/var/lib/postgresql/<major-version>/...` rather than directly in `/var/lib/postgresql/data`.
- If a host directory was previously created by older images at `/var/lib/postgresql/data`, mounting it directly to `/var/lib/postgresql/data` with the new image can cause layout conflicts and prevent proper upgrades.
- Recommended container mount: mount your host volume at `/var/lib/postgresql` (single top-level mount). This lets the container create versioned subdirectories and enables `pg_upgrade` without mount-boundary issues.

What changed and why
- Older images stored cluster data directly in `/var/lib/postgresql/data`.
- Newer images follow PostgreSQL's native layout and create per-major-version cluster directories (e.g. `/var/lib/postgresql/15/main` or `/var/lib/postgresql/18/main`).
- Mounting the host at `/var/lib/postgresql/data` prevents the image from creating the expected versioned directories and can lead to errors or require a manual `pg_upgrade`.

Quick fix for Docker Compose
- Change your Postgres service volume from:

```yaml
# (old)
- ./postgres-data:/var/lib/postgresql/data
```

to:

```yaml
# (recommended)
- ./postgres-data:/var/lib/postgresql
```

This gives the container a single mount point and allows the image to manage the versioned subdirectories.

Migration guidance (safe options)

1) If you do not have important data (fresh dev environment)
- Stop the container, change the compose mount to `/var/lib/postgresql`, and start the container. The image will initialize a new cluster under the mount.

2) If you have existing data created by an older image
- DO NOT simply start the new major-version image over the old data. Back up first.

Option A — Logical dump & restore (recommended, simplest):

```bash
# Dump everything from the old data (run using the old image if needed)
# Adjust container name, credentials, and old image tag (e.g., postgres:13)
docker run --rm -v "$PWD/postgres-data":/var/lib/postgresql/data -e POSTGRES_PASSWORD=root postgres:13 \
  bash -c "pg_ctl -D /var/lib/postgresql/data -w start && pg_dumpall -U root" > dump.sql

# After switching compose to mount at /var/lib/postgresql, start the new postgres image
docker compose up -d postgres

# Restore into the new running container (adjust container name)
cat dump.sql | docker exec -i <new-postgres-container-name> psql -U root -d postgres
```

Option B — Binary upgrade using `pg_upgrade` (advanced):
- This requires both old and new binaries and is more complex. Typical approach:
  1. Ensure the host mount point is `/var/lib/postgresql` so both images can access the same top-level directory.
  2. Run a container with the old postgres binary (old major version) and another with the new binary. Use `pg_upgrade` or the postgres image's documented approach to upgrade in-place.
  3. Test thoroughly and keep backups.

Note: `pg_upgrade` can be faster and preserve OIDs/large objects, but it requires care and ample testing. See the upstream docs and the docker-library/postgres examples.

Other recommendations
- Pin your Postgres image to a specific major version in `docker-compose.yml` (avoid `latest`) to keep upgrades explicit. Example: `image: postgres:18`.
- Always back up `./postgres-data` (copy the folder) before changing mounts or switching to a new major version.
- If you need help creating a `docker-compose` recipe to run an in-place `pg_upgrade` (old + new images), I can add an example compose and commands.

References
- Docker Library issue discussing layout: https://github.com/docker-library/postgres/pull/1259
- Long discussion & background: https://github.com/docker-library/postgres/issues/37

---
Created automatically: added guidance for Postgres 18+ volume mounts and safe migration steps.

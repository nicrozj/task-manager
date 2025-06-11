set -e

host="$POSTGRES_HOST"
db="$POSTGRES_DB"
user="$POSTGRES_USER"

echo "wait PostgreSQL ($host)..."

until pg_isready -h "$host" -p 5432 -U "$user" -d "$db"; do
  >&2 echo "Postgres is unvailable - waiting..."
  sleep 2
done

>&2 echo "Postgres is ready — start backend"

exec "$@"

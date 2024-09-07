#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until nc -z "$host" 3306; do
  echo "Waiting for MySQL to be ready..."
  sleep 1
done

echo "MySQL is ready! Executing command: $cmd"
exec $cmd
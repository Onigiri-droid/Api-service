#!/usr/bin/env bash
# Ожидает, пока указанный хост и порт станут доступными
host="$1"
shift
port="$1"
shift

while ! nc -z "$host" "$port"; do
  sleep 1
done

exec "$@"

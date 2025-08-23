#!/bin/sh
set -e

# Installe les deps si node_modules absent/incomplet
if [ ! -d "node_modules" ]; then
  echo ">> Installing node dependencies (npm ci)…"
  npm ci
fi

echo ">> Running: $@"
exec "$@"

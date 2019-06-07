#!/bin/bash -e

log() {
  echo -e "${NAMI_DEBUG:+${CYAN}${MODULE} ${MAGENTA}$(date "+%T.%2N ")}${RESET}${@}" >&2
}

setup_db() {
  log "Configuring the database"
  gorm-goose up
}

log "Waiting for Postgres..."
/root/wait-for-it.sh db:5432 --timeout=180 -- echo "PostgreSQL started"

setup_db

#log "Start cron..."
#service cron start

log "Start server"
go-api-boilerplate server

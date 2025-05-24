#!/bin/sh

set -e 
echo "" > .env
grep -v '^#' .env.example | grep -v '^$' | while IFS= read -r line; do
  envname="${line%%=*}"
  envdefaultvalue="${line#*=}"
  envcurrentvalue="$(eval echo \$$envname)"
  envvalue="${envcurrentvalue:-$envdefaultvalue}"

  echo "${envname}=${envvalue}" >> .env
done

echo ".env file generated successfully."

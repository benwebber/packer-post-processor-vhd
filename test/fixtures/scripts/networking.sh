#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Clear networking configuration.

rm -f /etc/udev/rules.d/70-persistent-net.rules

for dev in /etc/sysconfig/network-scripts/ifcfg-*; do
  if [[ "$(basename "${dev}")" != 'ifcfg-lo' ]]; then
    sed -i '/^HWADDR/d' "${dev}"
    sed -i '/^UUID/d' "${dev}"
  fi
done

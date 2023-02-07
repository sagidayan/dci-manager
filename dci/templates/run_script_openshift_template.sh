#!/usr/bin/env bash

echo "Executing DCI Job {{ .Name }}"
echo

set -x

# Ansible inside dci-openshift-app-agent-ctl uses it for colorful output.
export ANSIBLE_FORCE_COLOR=True

CONFIG={{ .JobFolder }}/settings.yml


dci-openshift-agent-ctl --config "${CONFIG}" -s

#!/bin/bash

# this file must be sourced
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))

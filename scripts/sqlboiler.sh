#!/bin/bash

readonly service="$1"

(cd "internal/${service}" && sqlboiler psql)

#!/bin/bash

# Load environment variables from .env.local if it exists
if [ -f .env.local ]; then
    echo "Loading environment variables from .env.local"
    set -a  # automatically export all variables
    source .env.local
    set +a  # disable automatic export
    echo "Environment variables loaded successfully"
else
    echo "No .env.local file found, using defaults"
fi

# Execute the command passed as arguments
exec "$@"

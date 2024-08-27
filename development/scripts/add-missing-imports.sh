#!/bin/bash

# List of files to process
files=("api_circuits.go" "api_dcim.go" "api_extras.go" "api_ipam.go" "api_tenancy.go" "api_users.go" "api_virtualization.go")

for file in "${files[@]}"
do
  full_path="/client/${file}"
  echo "Processing file: $file"
  
  # Check if the file contains "time" in the import block
  if awk '/import \(/,/\)/ { if ($0 ~ /"time"/) exit 1 }' "$full_path"; then
    echo "\"time\" package not imported in $file"
    # Add the "time" import to the import block if it exists
    if grep -q "import (" "$full_path"; then
      sed -i '/import (/a\'$'\t''"time"' "$full_path"
      echo "\"time\" package added to $file"
    else
      echo "No import block found in $file"
    fi
  else
    echo "\"time\" package already imported in $file"
  fi
done


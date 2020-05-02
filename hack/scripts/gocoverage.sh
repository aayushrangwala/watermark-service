#!/bin/bash
# This script will test the coverage of ../pkg and ../cmd
# If this script is called on a Presubmit, this will only test the coverage and exit 1 if coverage drops below 70%
# if Postsubmit, then it will upload the report to Codecov.io

set -ex

# NOTE: TODO: uncomment this and code below for coverage percentage check once we achieve stable stage
# EXPECTED_COVERAGE="70.0"
COVERAGE_PROFILE="coverage.out"

echo "Running go tool coverage from profile 'coverage.out'..."

make go-test

COVERAGE_PERCENTAGE=$(go tool cover -func=${COVERAGE_PROFILE}  | grep 'total:' | awk '{print $3}' | sed 's/%//')

if [[ -n "${PULL_NUMBER}" ]]; then
	# if (( ${COVERAGE_PERCENTAGE%%.*} < ${EXPECTED_COVERAGE%%.*} || ( ${COVERAGE_PERCENTAGE%%.*} == ${EXPECTED_COVERAGE%%.*} && ${COVERAGE_PERCENTAGE##*.} < ${EXPECTED_COVERAGE##*.} ) )) ; then
	# 	echo "coverage dropped to ${COVERAGE_PERCENTAGE}, expected was ${EXPECTED_COVERAGE}"
	# 	exit 1
	# fi
	echo "Coverage for ${PULL_NUMBER}:  ${COVERAGE_PERCENTAGE}"
else
	echo "Sending the report to CodeCov..."
	bash <(curl -s https://codecov.io/bash) -t "${CODECOV_REPO_TOKEN}"
fi

if [[ -n "${ARTIFACTS_DIR:-}" ]]; then
  	cp coverage.out "${ARTIFACTS_DIR}"
fi

exit 0

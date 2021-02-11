#!/usr/bin/env bash
# Copyright 2021 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.#!/usr/bin/env bash



set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
export KUBE_ROOT
source "${KUBE_ROOT}/hack/lib/init.sh"

migrated_files="${KUBE_ROOT}/hack/.structured_logging"
migrated_packages=$(sed "s| | -e |g" "${migrated_files}")

# Checking if any of the packages migrated to structured logging using unstructured method.
if grep -rE "klog.Info[^S]|klog.V.*\.Info[^S]|klog.Error[^S]|klog.Warning[^S]|klog.Fatal[^S]" $migrated_packages; then
	echo "This package is migrated to structured logging. Please see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/migration-to-structured-logging.md"
	exit 1
else
	echo "verified structured logging."
fi

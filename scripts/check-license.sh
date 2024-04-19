#!/bin/bash
#
# Copyright contributors to the Hyperledger Fabric Operator project
#
# SPDX-License-Identifier: Apache-2.0
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
#
# 	  http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

cat << EOB > golang_copyright.txt
/*
 * Copyright contributors to the Hyperledger Fabric Operator project
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 * 	  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

EOB

cat << EOB > shell_copyright.txt
#
# Copyright contributors to the Hyperledger Fabric Operator project
#
# SPDX-License-Identifier: Apache-2.0
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
#
# 	  http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

EOB

function filterGeneratedFiles {
    for f in $@; do
        head -n5 $f | grep -qE 'Code generated by.*DO NOT EDIT' || echo $f
    done
}

function filterExcludedFiles {
  CHECK=`echo "$CHECK" \
		| grep -v "^\.build/" \
		| grep -v "^\.git/" \
		| grep -v "^\.gitignore" \
		| grep -v "\.json$" \
		| grep -v "\.pem$" \
		| grep -v "\.crt$" \
		| grep -v "\.txt$" \
		| grep -v "\.md$" \
		| grep -v "_sk$" \
		| grep -v "\.key$" \
		| grep -v "\.gen\.go$" \
		| grep -v "tools/" \
		| grep -v "testdata/" \
		| grep -v "vendor/" \
		| grep -v "go.mod" \
		| grep -v "go.sum" \
		| grep -v .secrets.baseline \
		| grep -v .pre-commit-config.yaml \
		| grep -v .ibp.com_*.yaml \
		| grep -v .deepcopy.go \
		| sort -u`

  CHECK=$(filterGeneratedFiles "$CHECK")
}

CHECK=$(git diff --name-only --diff-filter=ACMRTUXB HEAD)
filterExcludedFiles
if [[ -z "$CHECK" ]]; then
  CHECK=$(git diff-tree --no-commit-id --name-only --diff-filter=ACMRTUXB -r "HEAD^..HEAD")
  filterExcludedFiles
fi

if [[ -z "$CHECK" ]]; then
   echo "All files are excluded from having license headers"
   exit 0
fi

missing=`echo "$CHECK" | xargs ls -d 2>/dev/null | xargs grep -L "SPDX-License-Identifier: Apache-2.0"`
if [[ -z "$missing" ]]; then
   echo "All files have SPDX-License-Identifier: Apache-2.0"
   exit 0
fi

TMPFILE="./tmpfile"

for FILE in ${missing}; do
	EXT="${FILE##*.}"
	echo "Adding copyright notice to $FILE"
	if [ "${EXT}" = "go" ]; then
		cat golang_copyright.txt ${FILE} > ${TMPFILE}
		cat ${TMPFILE} > ${FILE}
		rm -f ${TMPFILE}
		echo "	${FILE} copyright notice added"
	elif [ "${EXT}" = "yaml" ]; then
		cat shell_copyright.txt ${FILE} > ${TMPFILE}
		cat ${TMPFILE} > ${FILE}
		rm -f ${TMPFILE}
		echo "	${FILE} copyright notice added"
	elif [ "${EXT}" = "sh" ]; then
		cat shell_copyright.txt ${FILE} > ${TMPFILE}
		cat ${TMPFILE} > ${FILE}
		rm -f ${TMPFILE}
		echo "	${FILE} copyright notice added"
	else
		echo "invalid file extension"
	fi
done

rm golang_copyright.txt shell_copyright.txt

exit 0
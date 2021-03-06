#!/bin/bash

#
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# By default all sub-projects are built.  Supply a list of projects to build if
# only a subset is wanted.

# make sure we start out in traffic_control dir
topscript=$(readlink -f $0)
export TC_DIR=$(dirname $(dirname "$topscript"))
[[ -n $TC_DIR ]] && cd "$TC_DIR" || { echo "Could not cd $TC_DIR"; exit 1; }

. build/functions.sh

checkEnvironment


if [[ $# -gt 0 ]]; then
	projects=( "$*" )
else
	# get all subdirs containing build/build_rpm.sh
	projects_to_build=( */build/build_rpm.sh )
	# Always build tarball when building everything..
	projects=(tarball)
	for p in "${projects_to_build[@]}"; do
	  p=${p%%/*}
		if [[ $p != "traffic_monitor_golang" ]]; then
			projects+=($p)
		fi
	done
fi


declare -a badproj
declare -a goodproj
for p in "${projects[@]}"; do
	if [[ $p == tarball ]]; then
		if isInGitTree; then
			echo "-----  Building tarball ..."
			tarball=$(createTarball "$TC_DIR")
			ls -l $tarball
		else
			echo "---- Skipping tarball creation"
		fi
		continue
	fi
	if [[ $p == docs ]]; then
		if isInGitTree; then
			echo "-----  Building docs ..."
			pushd docs
			make html
			popd
			tarball=$(createDocsTarball "${TC_DIR}")
			ls -l $tarball
		else
			echo "---- Skipping docs creation"
		fi
		continue
	fi
	# strip trailing /
	p=${p%/}
	bldscript="$p/build/build_rpm.sh"
	if [[ ! -x $bldscript ]]; then
		echo "$bldscript not found"
		badproj+=($p)
		continue
	fi

	echo "-----  Building $p ..."
	if $bldscript; then
		goodproj+=($p)
	else
		echo "$p failed: $!"
		badproj+=($p)
	fi
done

if [[ ${#goodproj[@]} -ne 0 ]]; then
	echo "The following subdirectories built successfully: "
	for p in "${goodproj[@]}"; do
		echo "   $p"
	done
	echo "See $(pwd)/dist for newly built rpms."
fi

if [[ ${#badproj[@]} -ne 0 ]]; then
	echo "The following subdirectories had errors: "
	for p in "${badproj[@]}"; do
		echo "   $p"
	done
	exit 1
fi

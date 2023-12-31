#!/usr/bin/env bats

load helpers

TEST_CGROUP_NAME="runc-cgroups-integration-test"
CGROUP_MEMORY="${CGROUP_MEMORY_BASE_PATH}/${TEST_CGROUP_NAME}"

function teardown() {
    rm -f $BATS_TMPDIR/runc-update-integration-test.json
    teardown_running_container test_cgroups_kmem
    teardown_busybox
}

function setup() {
    teardown
    setup_busybox
}

function check_cgroup_value() {
    cgroup=$1
    source=$2
    expected=$3

    current=$(cat $cgroup/$source)
    echo  $cgroup/$source
    echo "current" $current "!?" "$expected"
    [ "$current" -eq "$expected" ]
}

@test "runc update --kernel-memory (initialized)" {
	# XXX: currently cgroups require root containers.
    requires cgroups_kmem root

    # Add cgroup path
    sed -i 's/\("linux": {\)/\1\n    "cgroupsPath": "\/runc-cgroups-integration-test",/'  ${BUSYBOX_BUNDLE}/config.json

    # Set some initial known values
    DATA=$(cat <<-EOF
    "memory": {
        "kernel": 16777216
    },
EOF
    )
    DATA=$(echo ${DATA} | sed 's/\n/\\n/g')
    sed -i "s/\(\"resources\": {\)/\1\n${DATA}/" ${BUSYBOX_BUNDLE}/config.json

    # run a detached busybox to work with
    runc run -d --console-socket $CONSOLE_SOCKET test_cgroups_kmem
    [ "$status" -eq 0 ]

    # update kernel memory limit
    runc update test_cgroups_kmem --kernel-memory 50331648
    [ "$status" -eq 0 ]

	# check the value
    check_cgroup_value $CGROUP_MEMORY "memory.kmem.limit_in_bytes" 50331648
}

@test "runc update --kernel-memory (uninitialized)" {
	# XXX: currently cgroups require root containers.
    requires cgroups_kmem root

    # Add cgroup path
    sed -i 's/\("linux": {\)/\1\n    "cgroupsPath": "\/runc-cgroups-integration-test",/'  ${BUSYBOX_BUNDLE}/config.json

    # run a detached busybox to work with
    runc run -d --console-socket $CONSOLE_SOCKET test_cgroups_kmem
    [ "$status" -eq 0 ]

    # update kernel memory limit
    runc update test_cgroups_kmem --kernel-memory 50331648
    # Since kernel 4.6, we can update kernel memory without initialization
    # because it's accounted by default.
    if [ "$KERNEL_MAJOR" -lt 4 ] || [ "$KERNEL_MAJOR" -eq 4 -a "$KERNEL_MINOR" -le 5 ]; then
        [ ! "$status" -eq 0 ]
    else
        [ "$status" -eq 0 ]
        check_cgroup_value $CGROUP_MEMORY "memory.kmem.limit_in_bytes" 50331648
    fi
}

@test "runc create (rootless + no limits + no cgrouppath + no permission) succeeds" {
   requires rootless

    runc run -d --console-socket $CONSOLE_SOCKET test_rootless_cgroups
    [ "$status" -eq 0 ]
}

@test "runc create (rootless + no limits + cgrouppath + no permission) fails with permission error" {
   requires rootless

    # Add cgroup path
    sed -i 's/\("linux": {\)/\1\n    "cgroupsPath": "\/runc-cgroups-integration-test",/'  ${BUSYBOX_BUNDLE}/config.json

    runc run -d --console-socket $CONSOLE_SOCKET test_rootless_cgroups
    [ "$status" -eq 1 ]
    [[ ${lines[1]} == *"permission denied"* ]]
}

@test "runc create (rootless + limits + no cgrouppath + no permission) fails with informative error" {
   requires rootless

    # Add resource limit
    sed -i 's/\("linux": {\)/\1\n   "resources": { "pids": { "limit": 100 } },/'  ${BUSYBOX_BUNDLE}/config.json

    runc run -d --console-socket $CONSOLE_SOCKET test_rootless_cgroups
    [ "$status" -eq 1 ]
    [[ ${lines[1]} == *"cannot set limits on the pids cgroup, as the container has not joined it. The cgroup controller may not be mounted or you may not have previously had permissions to create the cgroup subdirectory"* ]]
}

@test "runc create (rootless + limits + cgrouppath + permission on the cgroup dir) succeeds" {
   requires rootless
   requires rootless_cgroups
   # this test expects that the /sys/fs/cgroup/*/rootless_cgroup directories have already been created
   # and chowned to the rootless user.

    # Add cgroup path
    sed -i 's/\("linux": {\)/\1\n    "cgroupsPath": "\/rootless_cgroup\/test_rootless_cgroups",/'  ${BUSYBOX_BUNDLE}/config.json

    # Add resource limit
    sed -i 's/\("linux": {\)/\1\n   "resources": { "pids": { "limit": 100 } },/'  ${BUSYBOX_BUNDLE}/config.json

    runc run -d --console-socket $CONSOLE_SOCKET test_rootless_cgroups
    [ "$status" -eq 0 ]
}

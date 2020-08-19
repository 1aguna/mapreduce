package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"net/rpc"
	"strconv"
)

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.
type JobType string

type JobArgs struct {
	 File string
	 Operation 	JobType
	 JobID int
	 NumJobsOther int // number of jobs in other phases (map or reduce)
}

type JobReply struct {
	OK bool
}

type RegisterArgs struct {
}

type RegisterReply struct {
	workerID int
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the master.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func masterSock() string {
	s := "/var/tmp/824-mr_"
	s += strconv.Itoa(os.Getuid())
	return s
}

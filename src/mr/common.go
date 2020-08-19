package mr

const (
	MAP 	= 1
	REDUCE 	= 0
)

type Job struct {
	Filename 	string
	NReduce 	int
	NMaps 		int
	Seq 		int
	Phase		int
	Alive 		bool
}


package pi

import "errors"

const (
	WordDir    = "./tmp"
	PipeRead   = "./tmp/pipe.1"
	PipeWrite  = "./tmp/pipe.2"
	PythonPath = "/usr/bin/python3"
	ScriptPath = "./script/raspberry.py"
)

// parking
const (
	EmptyParkingSpace    int = 0
	NonEmptyParkingSpace int = 1

	FrontCamera int = 0
	RearCamera  int = 1
)

// parkingMgr
var (
	NoErr                   error = nil
	ErrNoParkingSpace       error = errors.New("no parking space")
	ErrLicenseAlreadyExists error = errors.New("the license plate already exists")
	ErrLicenseNotExists     error = errors.New("the license plate does not exist")
)

// operate
var (
	ErrProtocol         error = errors.New("protocol error")
	ErrArrayOutOfBounds error = errors.New("array out of bounds")
)

package pi

import "errors"

const (
	WordDir    = "./tmp"
	PipeRead   = "./tmp/pipe.1"
	PipeWrite  = "./tmp/pipe.2"
	PythonPath = "/usr/bin/python3"
	ScriptPath = "./script/raspberry.py"
)

//parking
var (
	EmptyParkingSpace    int = 0
	NonEmptyParkingSpace int = 1
)

//parkingMgr
var (
	NoErr                   error = nil
	ErrNoParkingSpace       error = errors.New("no parking space")
	ErrLicenseAlreadyExists error = errors.New("the license plate already exists")
	ErrLicenseNotExists     error = errors.New("the license plate does not exist")
)

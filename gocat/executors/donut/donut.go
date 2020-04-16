// +build windows

package donut


import (
	"io/ioutil"
	"runtime"

	"github.com/mitre/sandcat/gocat/executors/execute"
)

type Donut struct {
	archName string
}

func init() {
	runner := &Donut{
		archName: "donut_"+runtime.GOARCH,
	}
	if runner.CheckIfAvailable() {
		execute.Executors[runner.archName] = runner
	}
}

func (d *Donut) Run(command string, timeout int) ([]byte, string, string) {
	bytes, _ := ioutil.ReadFile("something.donut")

	// Create the new host process, getting its output

	//Pass the shellcode bytes and PID of the host process to the Runner



	//gToByteArrayString(string(content))

	//bytes, _ := util.StringToByteArrayString(command)
	//task, pid := Runner(bytes)

	//return status and also output of process

	task, pid := Runner(bytes)
	if task {
		return []byte("Shellcode executed successfully."), execute.SUCCESS_STATUS, pid
	}
	return []byte("Shellcode execution failed."), execute.ERROR_STATUS, pid
}

func (d *Donut) String() string {
	return d.archName
}

func (d *Donut) CheckIfAvailable() bool {
	return IsAvailable()
}

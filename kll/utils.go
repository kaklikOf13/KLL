package kll

import (
	"fmt"
	"os"
	"path"
	"syscall"
	"unsafe"
)

type Callstack struct {
	Line uint
	Col  uint
	Show string
}

func containsString(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func splitWithSeparators(input []Token, separators []string) []Block {
	result := make([]Block, 0, 13)
	currentSegment := make(Block, 0, 6)

	for _, item := range input {
		containsSeparator := false

		for _, separator := range separators {
			if item.Type == separator {
				containsSeparator = true
				break
			}
		}

		if containsSeparator {
			if len(currentSegment) > 0 {
				result = append(result, currentSegment)
				currentSegment = make([]Token, 0)
				continue
			}
		}

		currentSegment = append(currentSegment, item)
	}

	if len(currentSegment) > 0 {
		result = append(result, currentSegment)
	}

	return result
}

func WriteFile(filename string, content []byte) error {
	// Create the file
	os.MkdirAll(path.Dir(filename), os.ModePerm)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the bytes to the file
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func GetJIT[JITFunc any](bytecode []byte) (JITFunc, error) {
	var ret JITFunc
	pagesize := syscall.Getpagesize()
	memSize := (len(bytecode) + pagesize - 1) &^ (pagesize - 1)
	mem, err := syscall.Mmap(-1, 0, memSize, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_ANON|syscall.MAP_SHARED)
	if err != nil {
		return ret, fmt.Errorf("memory allocation failed: %v", err)
	}
	copy(mem, bytecode)
	*(*uintptr)(unsafe.Pointer(&ret)) = (uintptr)(unsafe.Pointer(&mem))
	return ret, nil
}

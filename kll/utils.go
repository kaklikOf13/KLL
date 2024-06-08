package kll

import (
	"bytes"
	"encoding/binary"
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

// IntToBytes converts an int to a byte slice.
func Int64ToBytes(n int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(n))
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Int32ToBytes converts an int32 to a byte slice.
func Int32ToBytes(n int32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, n)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Int16ToBytes converts an int16 to a byte slice.
func Int16ToBytes(n int16) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, n)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Float32ToBytes converts a float32 to a byte slice.
func Float32ToBytes(f float32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Float64ToBytes converts a float64 to a byte slice.
func Float64ToBytes(f float64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// Package libtextcat is a Go wrapper for C library 'libtextcat'.
package libtextcat

// #cgo LDFLAGS: -ltextcat
// #include <stdlib.h>
// #include "textcat.h"
import "C"
import (
    "sync"
    "fmt"
    "unsafe"
    "strings"
)

// TextCat wrapper provides libtextcat functionality.
type TextCat struct {
    cat        unsafe.Pointer  // Pointer to data used by libtextcat. Must be freed after allocation.
    mutex      sync.Mutex 
}

// NewTextCat creates a new libtextcat wrapper that can be used to classify text. If it is successfully created, it
// must be closed as it needs to free native resources.
func NewTextCat(configPath string) (*TextCat, error) {
    tc := new(TextCat)

    configStr := C.CString(configPath)
    defer C.free(unsafe.Pointer(configStr))

    tcPointer := C.textcat_Init(configStr);
    if tcPointer == nil {
        return nil, fmt.Errorf("Cannot create textcat wrapper with config file: '%s'", configStr)
    }

    tc.cat = tcPointer

    return tc, nil
}

// Classify returns classification result
func (tc *TextCat) Classify(text string) ([]string, error) {
    textStr := C.CString(text)
    defer C.free(unsafe.Pointer(textStr))

    cResult := C.textcat_Classify(tc.cat, 
                                  textStr, 
                                  (C.size_t)(len(text)))

    if cResult == nil {
        return nil, fmt.Errorf("Cannot classify %s", text)
    }

    goResult := C.GoString(cResult)

    // Original libtextcat library returns result in the following format:
    // [match1][match2][match3]...
    //
    // So, to get the result slice, we just cut off first and last braces and
    // then split the string by "]["

    if len(goResult) < 2 {
        return nil, fmt.Errorf("Classify returned invalid result")
    }

    goResult = goResult[1 : len(goResult) - 1]
    return strings.Split(goResult, "]["), nil
}

// Close frees native C resources
func (tc *TextCat) Close() {
    tc.mutex.Lock()
    defer tc.mutex.Unlock()

    if tc.cat != nil {
        C.textcat_Done(tc.cat)
    }
}

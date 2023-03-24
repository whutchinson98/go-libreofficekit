package libreofficekit

/*
#cgo CFLAGS: -I ./ -D LOK_USE_UNSTABLE_API
#cgo LDFLAGS: -ldl
#include <lokbridge.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

type Office struct {
	handle *C.struct__LibreOfficeKit
	Mutex  *sync.Mutex
}

// NewOffice returns new Office or error if LibreOfficeKit fails to load
// required libs (actually, when libreofficekit-dev package isn't installed or path is invalid)
func NewOffice(path string) (*Office, error) {
	office := new(Office)

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	lokit := C.lok_init(cPath)
	if lokit == nil {
		return nil, fmt.Errorf("failed to initialize LibreOfficeKit with path: '%s'", path)
	}

	office.handle = lokit
	office.Mutex = &sync.Mutex{}

	return office, nil

}

// Close destroys C LibreOfficeKit instance
func (office *Office) Close() {
	C.destroy_office(office.handle)
}

// GetError returns last happened error message in human-readable format
func (office *Office) GetError() string {
	message := C.get_error(office.handle)
	return C.GoString(message)
}

// LoadDocument return Document or error, if LibreOffice fails to open document at provided path.
// Actual error message can be retrieved by office.GetError method
func (office *Office) LoadDocument(path string) (*Document, error) {
	document := new(Document)
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	handle := C.document_load(office.handle, cPath)
	if handle == nil {
		return nil, fmt.Errorf("failed to load document")
	}
	document.handle = handle
	return document, nil
}

type Document struct {
	handle *C.struct__LibreOfficeKitDocument
}

// Close destroys document
func (document *Document) Close() {
	C.destroy_document(document.handle)
}

// SaveAs saves document at desired path in desired format with applied filter rules
// Actual (from libreoffice) error message can be read with Office.GetError
func (document *Document) SaveAs(path string, format string, filter string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	cFilter := C.CString(filter)
	defer C.free(unsafe.Pointer(cFilter))
	status := C.document_save(document.handle, cPath, cFormat, cFilter)
	if status != 1 {
		return fmt.Errorf("failed to save document")
	}
	return nil
}

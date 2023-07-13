#include <stdlib.h>
#include "./LibreOfficeKitInit.h"
#include "./LibreOfficeKit.h"


void destroy_office(LibreOfficeKit* pThis) {
    return pThis->pClass->destroy(pThis);
};

char* get_error(LibreOfficeKit* pThis) {
    return pThis->pClass->getError(pThis);
};

void free_error(LibreOfficeKit* pThis, char* message) {
    return pThis->pClass->freeError(message);
}

LibreOfficeKitDocument* document_load(LibreOfficeKit* pThis, const char* pURL) {
    return pThis->pClass->documentLoad(pThis, pURL);
};

void destroy_document(LibreOfficeKitDocument* pThis) {
    return pThis->pClass->destroy(pThis);
};

int document_save(LibreOfficeKitDocument* pThis, const char* pUrl, const char* pFormat, const char* pFilterOptions) {
    return pThis->pClass->saveAs(pThis, pUrl, pFormat, pFilterOptions);
};

void post_uno_command(LibreOfficeKitDocument* pThis, const char* pCommand, const char* pArguments, bool bNotifyWhenFinished) {
    return pThis->pClass->postUnoCommand(pThis, pCommand, pArguments, bNotifyWhenFinished);
}

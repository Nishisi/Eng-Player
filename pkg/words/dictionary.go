package words

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreServices

#import <Foundation/Foundation.h>
#import <CoreServices/CoreServices.h>

const char* dictionary(char *searchword) {
  NSString* word = [NSString stringWithUTF8String:searchword];
  return [(NSString*)DCSCopyTextDefinition(NULL, (CFStringRef)word, CFRangeMake(0, [word length])) UTF8String];
  }
*/
import "C"
import "unsafe"

// 辞書データを取ってきます
func getDictionary(word string) string {
	w := C.CString(word)
	defer C.free(unsafe.Pointer(w))

	result := C.dictionary(w)
	return C.GoString(result)
}

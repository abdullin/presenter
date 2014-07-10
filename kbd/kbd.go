package kbd

/*
#cgo CFLAGS: -Qunused-arguments
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>
#include <CoreFoundation/CoreFoundation.h>
#include <Carbon/Carbon.h>

void keyevt(int keycode, bool isdown) {
    CGEventRef evt;
    evt = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)keycode, isdown);
    CGEventPost(kCGSessionEventTap, evt);
    CFRelease(evt);
}
*/
import "C"

var KEYS = map[string]int32{
	"KEY_LEFT":  C.kVK_LeftArrow,
	"KEY_RIGHT": C.kVK_RightArrow,
}

func KeyPress(key string) {
	C.keyevt(C.int(KEYS[key]), C.bool(true))
	C.keyevt(C.int(KEYS[key]), C.bool(false))
}

package main
//#cgo amd64 386 CFLAGS: -DX86_64=1 -arch x86_64  -Wall -fPIC
//#cgo CFLAGS: -I${SRCDIR}/library/include
//#cgo LDFLAGS: -L${SRCDIR}/library/lib -lavcodec -lavdevice -lavformat -lavfilter -lavutil  -lswscale -lswresample
//#cgo LDFLAGS: -framework CoreAudio -framework AudioToolbox -framework AudioUnit -framework Carbon -framework CoreMedia
//#cgo LDFLAGS: -framework MediaToolbox
//#cgo LDFLAGS: -framework CoreGraphics -framework VideoToolBox -liconv -framework Accelerate -lbz2
//#cgo LDFLAGS: -Wl,-framework,CoreFoundation -Wl,-framework,Security -Wl,-framework,VideoToolbox -Wl,-framework,CoreMedia -Wl,-framework,CoreVideo
//#include "hello.c"
import "C"
import "fmt"

func main() {
	Hello()
	fmt.Println("hello 123")
}

func Hello() error {
	_, err := C.hello_world()
	return err
}

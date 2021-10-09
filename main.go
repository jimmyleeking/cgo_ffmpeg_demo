package main

//#cgo amd64 386 CFLAGS: -DX86_64=1 -arch x86_64  -Wall -fPIC
//#cgo CFLAGS: -I${SRCDIR}/library/include
//#cgo LDFLAGS: -L${SRCDIR}/library/lib -laom -lavif -ljpeg -lturbojpeg -lyuv
//#include "iccjpeg.c"
//#include "avifutil.c"
//#include "y4m.c"
//#include "avifjpeg.c"
//#include "bridge.c"
import "C"
import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

func main() {
	start := time.Now().Unix()
	totalTime := 100
	group := sync.WaitGroup{}
	for i := 0; i < totalTime; i++ {
		group.Add(1)
		go func(tindex int) {

			index := strconv.Itoa(tindex)
			err := avif2jpeg("./res/sample.avif", "./testresult/nice"+index+".jpeg")
			if err != nil {
				return
			}
			group.Done()
		}(i)
	}
	group.Wait()
	span := time.Now().Unix() - start
	fmt.Println("time:", totalTime, ",usage time:", span)

}

func avif2jpeg(inputFile string, outputFile string) error {
	input := C.CString(inputFile)
	output := C.CString(outputFile)
	_, err := C.avif2jpeg(input, output)
	C.free(unsafe.Pointer(input))
	C.free(unsafe.Pointer(output))
	return err
}

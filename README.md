# Demo for Cgo With ffmpeg

This is a demo for use ffmpeg by Golang project.

It would be so simple to assembly C library by cgo .


>It's only build ffmpeg for osx_64.if you want to support other platform,replace other platform library and edit the cgo import link flag)

project tree list:

```
├── go.mod
├── library
│   ├── include
│   │   ├── hello.c // own c bridge code
│   │   ├── libavcodec
│   │   ├── libavdevice
│   │   ├── libavfilter
│   │   ├── libavformat
│   │   ├── libavutil
│   │   ├── libswresample
│   │   └── libswscale
│   └── lib
│       ├── libavcodec.a
│       ├── libavdevice.a
│       ├── libavfilter.a
│       ├── libavformat.a
│       ├── libavutil.a
│       ├── libswresample.a
│       ├── libswscale.a
│       └── pkgconfig
└── main.go // main function

```

Running demo by command:

```azure
go build main.go
./main
```
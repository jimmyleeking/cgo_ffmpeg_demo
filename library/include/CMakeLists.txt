cmake_minimum_required(VERSION 3.5)

project(test)

include_directories("./include/")
link_directories("../lib")

add_executable(test bridge.c avifjpeg.c avifutil.c iccjpeg.c y4m.c)
target_link_libraries(test libaom.a libavif.a libjpeg.a libturbojpeg.a libyuv.a)
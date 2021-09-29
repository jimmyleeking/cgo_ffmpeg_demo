#include <stdio.h>
#include <libavformat/avformat.h>

void hello_world(){
    printf("version:%d",avformat_version());
}
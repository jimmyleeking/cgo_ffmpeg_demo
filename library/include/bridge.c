// Copyright 2019 Joe Drago. All rights reserved.
// SPDX-License-Identifier: BSD-2-Clause

#include "avif/avif.h"

#include "avifjpeg.h"
#include "avifutil.h"
#include "y4m.h"

#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DEFAULT_JPEG_QUALITY 90


/**
 * 转换数据到Jpg文件
 * @param inputFilename 输入文件地址
 * @param outputFilename 输出文件地址
 * @return
 */
int avif2jpeg(char *inputFilename, char *outputFilename) {

    printf("input file :%s ,output file:%s \n ",inputFilename,outputFilename);
    int jobs = 1;
    int jpegQuality = DEFAULT_JPEG_QUALITY;
    avifCodecChoice codecChoice = AVIF_CODEC_CHOICE_AUTO;
    avifChromaUpsampling chromaUpsampling = AVIF_CHROMA_UPSAMPLING_AUTOMATIC;
    avifStrictFlags strictFlags = AVIF_STRICT_ENABLED;

    if (!inputFilename) {
        return 1;
    }
    int returnCode = 0;
    avifImage *avif = avifImageCreateEmpty();
    avifDecoder *decoder = avifDecoderCreate();
    decoder->maxThreads = jobs;
    decoder->codecChoice = codecChoice;
    decoder->strictFlags = strictFlags;
    avifResult decodeResult = avifDecoderReadFile(decoder, avif, inputFilename);
    if (decodeResult == AVIF_RESULT_OK) {

        if (!avifJPEGWrite(outputFilename, avif, jpegQuality, chromaUpsampling)) {
            returnCode = 1;
        }
    } else {

        avifDumpDiagnostics(&decoder->diag);
        returnCode = 1;
    }
    avifDecoderDestroy(decoder);
    avifImageDestroy(avif);
    return returnCode;
}

void test(){

    if (AVIF_RESULT_OK == avif2jpeg("../../../res/sample.avif", "good3.jpg")) {
        printf("success");
    } else {
        printf("fail");
    }
}

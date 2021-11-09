#!/bin/bash
SRC_DIR="../../.."
DST_DIR="../../../pb_src"
./protoc -I=$SRC_DIR --cpp_out=$DST_DIR $SRC_DIR/test.proto

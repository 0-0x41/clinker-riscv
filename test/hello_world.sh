#!/bin/bash

test_name=$(basename "$0" .sh)
t=out/test/$test_name

mkdir -p "$t"

cat <<EOF | riscv64-linux-gnu-gcc -o "$t"/a.o -c -xc -
#include <stdio.h>

int main(void)
{
    printf("Hello World!");
    return 0;
}
EOF

./clinker-riscv "$t"/a.o

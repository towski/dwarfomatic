#define LINUX_BUILD
#include <iostream>
#include <stdint.h>
#include "foo.hpp"
#include "DataDefs.h"
#include "Export.h"
#include "item.h"
#include "unit.h"
//#include "library/include/df/item.h"
#include <sys/types.h>
#include <sys/mman.h>
#include <err.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void cxxFoo::Bar(void) {
    df::item *item;
    df::unit *unit;
    int fd = -1;
    char *zero;
    if ((fd = open("/tmp/test.dat", O_RDWR, 0)) == -1)
            err(1, "open");
    //anon = (char*)mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_ANON|MAP_SHARED, -1, 0);
    zero = (char*)mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_FILE|MAP_SHARED, fd, 0);
    if (zero == MAP_FAILED)
            errx(1, "either mmap");
    const char src[50] = "http://www.tutorialspoint.com";
    unit = (df::unit*)zero;
    //memcpy(&item, src, 30);
    //munmap(zero, 4096);
    //close(fd);

    std::cout<< unit->sex;
    std::cout<< "sizeof" << sizeof(df::unit);
	std::cout<<this->a<<std::endl;
}


#include <sys/types.h>
#include <sys/mman.h>
#include <err.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
 
/* Does not work on OSX, as you can't mmap over /dev/zero */
int main(void)
{
        const char str1[] = "string 1";
        const char str2[] = "string 2";
        int parpid = getpid(), childpid;
        int fd = -1;
        char *zero;
 
        if ((fd = open("/tmp/test.dat", O_RDWR, 0)) == -1)
                err(1, "open");
 
        //anon = (char*)mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_ANON|MAP_SHARED, -1, 0);
        zero = (char*)mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_FILE|MAP_SHARED, fd, 0);
 
        if (zero == MAP_FAILED)
                errx(1, "either mmap");
 
        strcpy(zero, str1);
 
        printf("PID %d:\t zero-backed %s\n", parpid, zero);
        switch ((childpid = fork())) {
        case -1:
                err(1, "fork");
                /* NOTREACHED */
 
        case 0:
                childpid = getpid();
                printf("PID %d:\t zero-backed %s\n", childpid, zero);
                sleep(3);
 
                printf("PID %d:\t zero-backed %s\n", childpid, zero);
                munmap(zero, 4096);
                close(fd);
                return (EXIT_SUCCESS);
        }
 
        sleep(2);
        strcpy(zero, str2);
 
        printf("PID %d:\tzero-backed %s\n", parpid, zero);
        munmap(zero, 4096);
        close(fd);
        return (EXIT_SUCCESS);
}

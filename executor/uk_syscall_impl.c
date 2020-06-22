#include <fcntl.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/ioctl.h>
#include <sys/mman.h>
#include <sys/prctl.h>
#include <sys/syscall.h>
#include <unistd.h>

#include <dlfcn.h>
#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>

static int (*sys_uk_hexdumpsn)(char *, size_t, const void *, size_t, size_t, int, unsigned int, const char *);

intptr_t execute_uk_hexdumpsn(intptr_t a[8]) {
	intptr_t res;
	void *lib;

	lib = dlopen("/root/out-rdynamic", RTLD_NOW);
	if (!lib) {
		fprintf(stderr, "dlopen %s\n", dlerror());
		errno = EINVAL;
	}

	sys_uk_hexdumpsn = (int (*)(char *, size_t, const void *, size_t, size_t, int, unsigned int, const char *)) dlsym(lib, "uk_hexdumpsn");
	if (!sys_uk_hexdumpsn) {
		fprintf(stderr, "dlsym %s\n", dlerror());
		errno = EINVAL;
	}

	res = sys_uk_hexdumpsn((char *) a[0], a[1], (const void *) a[2], a[3], a[4], a[5], a[6], (const char *) a[7]);

	return res;
}

static int (*sys_strlen)(const char *);

intptr_t execute_strlen(intptr_t a[8]) {
	intptr_t res;
	void *lib;

	lib = dlopen("/root/out-rdynamic", RTLD_NOW);
	if (!lib) {
		fprintf(stderr, "dlopen %s\n", dlerror());
		errno = EINVAL;
	}

	sys_strlen = (int (*)(const char *)) dlsym(lib, "strlen");
	if (!sys_strlen) {
		fprintf(stderr, "dlsym %s\n", dlerror());
		errno = EINVAL;
	}

	res = sys_strlen((const char *)a[0]);

	return res;
}
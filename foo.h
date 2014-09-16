#ifndef _MY_PACKAGE_FOO_H_
#define _MY_PACKAGE_FOO_H_

#ifdef __cplusplus
extern "C" {
#endif

	typedef void* Foo;
	Foo FooInit(void);
	void FooFree(Foo);
    int Size();
	const char* GetFirstName(int);
	const char* GetThoughts(int);
	const char* GetCurrentJob(int);

#ifdef __cplusplus
}
#endif

#endif

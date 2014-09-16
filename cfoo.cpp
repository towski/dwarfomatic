#include "foo.hpp"
#include "foo.h"

cxxFoo * foo;
Foo FooInit() {
	foo = new cxxFoo(1);
    foo->Init();
	return (void*)foo;
}

void FooFree(Foo f) {
	cxxFoo * foo = (cxxFoo*)f;
	delete foo;
}

int Size() {
	foo->Size();
}

const char* GetFirstName(int i) {
	foo->GetFirstName(i);
}

const char* GetThoughts(int i) {
	foo->GetThoughts(i);
}

const char* GetCurrentJob(int i) {
	foo->GetCurrentJob(i);
}

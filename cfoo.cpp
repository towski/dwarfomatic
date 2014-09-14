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

const char* FooBar(int i) {
	foo->Bar(i);
}

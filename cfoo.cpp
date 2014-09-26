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

const char* GetJobType(int i) {
	foo->GetJobType(i);
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

const char* GetHappiness(int i) {
	foo->GetHappiness(i);
}

int GetId(int i) {
	foo->GetId(i);
}

void Update() {
    foo->Update();
}

void Exit() {
    foo->Exit();
}

int Paused() {
    foo->Paused();
}

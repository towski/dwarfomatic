#include "foo.cpp"
#include "foo.h"
#include "foo.hpp"

int main(){
    cxxFoo *foo = new cxxFoo(1);
    //FooBar(cxxFoo);
    foo->Bar();
}

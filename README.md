# Wrapping Dwarf Fortress C++ in go

Requires go to be running in 32-bit, haven't tried it with 64-bit

First
go get github.com/towski/artery 
build_go.sh will build it and install it to .
./artery will start the web application

Then build the exporter with make
./runfoo.sh will export the data to ./public/

The templates for building the files are in ./templates

I borrowed most of the code from [a stack overflow answer](http://stackoverflow.com/questions/1713214/how-to-use-c-in-go), but modified it to use static linking.


# Wrapping Dwarf Fortress C++ in go

build_go.sh will build the web application, which lives at /df/artery
Build the binary with make

Then ./runfoo.sh will run the data exporter to ./public/

The templates for building the files are in ./templates

I borrowed most of the code from [a stack overflow answer](http://stackoverflow.com/questions/1713214/how-to-use-c-in-go), but modified it to use static linking.


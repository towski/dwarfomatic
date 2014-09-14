.PHONY: clean

TARGET=how

$(TARGET): libfoo.a
	~/save/go/bin/go build foo.go

libfoo.a: foo.o cfoo.o  #_obj/_cgo_.o
	ar r $@ $^

%.o: %.cpp
	#g++ -I library/include/modules/ -I library/include -I library/proto -O2 -o $@ -c $^
	g++ -m32 -I library/include/df -I library/include/ -I library/proto -I depends/protobuf/google/protobuf/ -I depends/protobuf -o $@ -c $^ -std=c++11
main:
	g++ -m32 -I library/include/df -I library/include/ -I library/proto -I depends/protobuf/google/protobuf/ -I depends/protobuf  \
	-Wl,/home/towski/code/howto-go-with-cpp/libfoo.a,/home/towski/save/df_linux/hack/libprotobuf-lite.so,/home/towski/save/df_linux/hack/libdfhack-client.so \
	main.cpp -std=c++11

clean:
	rm -f *.o *.so *.a $(TARGET)

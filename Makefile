.PHONY: clean

TARGET=how

$(TARGET): cpp/libfoo.a
	~/save/go/bin/go build cmd/client/client.go

cpp/libfoo.a: cpp/foo.o cpp/cfoo.o  #_obj/_cgo_.o
	cd cpp
	ar r $@ $^
	cd ..

cpp/libfoo.so: cpp/foo.o cpp/cfoo.o
	cd cpp
	gcc -m32 -shared -o libfoo.so foo.o cfoo.o /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++ -std=c++11
	cd ..

cpp/foo.o: cpp/foo.cpp
	cd cpp
	#g++ -I library/include/modules/ -I library/include -I library/proto -O2 -o $@ -c $^
	g++ -m32 -I library/include/df -I library/include/ -I library/proto -I depends/protobuf/google/protobuf/ -I depends/protobuf -o $@ -c $^ -std=c++11
	cd ..

cpp/cfoo.o: cpp/cfoo.cpp
	cd cpp
	g++ -m32 -I library/include/df -I library/include/ -I library/proto -I depends/protobuf/google/protobuf/ -I depends/protobuf -o $@ -c $^ -std=c++11
	cd ..

main: 
	g++ -m32 -I library/include/df -I library/include/ -I library/proto -I depends/protobuf/google/protobuf/ -I depends/protobuf  \
	-Wl,/home/towski/save/df_linux/hack/liblua.so,/home/towski/save/df_linux/hack/libprotobuf-lite.so,/home/towski/save/df_linux/hack/libdfhack.so,/home/towski/save/df_linux/hack/plugins/workflow.plug.so \
	foo.cpp -lstdc++ -std=c++11 

clean:
	rm -f cpp/*.o cpp/*.so cpp/*.a $(TARGET)

	#-Wl,/home/towski/save/df_linux/hack/libprotobuf-lite.so,/home/towski/save/df_linux/hack/libdfhack.so \

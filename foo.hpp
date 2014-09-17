#ifndef _MY_PACKAGE_FOO_HPP_
#define _MY_PACKAGE_FOO_HPP_

class cxxFoo {
public:
	int a;
	cxxFoo(int _a):a(_a){};
	~cxxFoo(){};
	int Size();
	const char* GetJobType(int);
	const char* GetFirstName(int);
	const char* GetThoughts(int);
	const char* GetHappiness(int);
	const char* GetCurrentJob(int);
	const char* Gender(int);
	void Init();
};

#endif

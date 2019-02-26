#include <limits.h>

int reverse(int x) {

	int Source = x;
	int Result = 0;
	while(Source != 0)
	{
		if(Result > INT_MAX/10)
			return 0;
		if(Result < INT_MIN/10)
			return 0;
		Result = (Result * 10) + (Source % 10);
		Source /= 10;
	}
	return Result;
}

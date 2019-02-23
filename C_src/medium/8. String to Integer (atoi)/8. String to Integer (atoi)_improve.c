#include<limits.h>
int myAtoi(char* str) {
	if(((*str <'0') || (*str > '9')) && ((*str != ' ') && (*str != '-') && (*str != '+')))
	{
		return 0;
	}

	while(1) // dealing white space
	{
		if(*str != ' ')
		{
			break;
		}
		str++;
	}

	unsigned char bSigned = 0;
	if(*str == '-')
	{
		bSigned = 1;
		str++;
	}
	else if(*str == '+')
	{
		str++;
	}

	int s32Sum =0;
	//parsing digital
	while(1)
	{
		if((*str < '0') || (*str > '9'))
		{
			return bSigned? -s32Sum : s32Sum;
		}
		if(s32Sum > (INT_MAX - (*str - '0')) / 10)
		{
			if(bSigned)
				return INT_MIN;
			else
				return INT_MAX;
		}
		s32Sum = (s32Sum * 10) + (*str - '0');
		str++;
	}
}


#include<limits.h>
#include<stdio.h>
int myAtoi(char* str) {
	if(((*str <'0') || (*str > '9')) && ((*str != ' ') && (*str != '-') && (*str != '+')))
	{
		printf("Is word %c \n", *str);
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
	long int s64CheckOverflow = 0;
	int s32Temp = 0;
	//parsing digital
	while(1)
	{
		printf("s32Sum = %d\n", s32Sum);
		if((*str < '0') || (*str > '9'))
		{
			return bSigned? -s32Sum : s32Sum;
		}
		s64CheckOverflow = (long int)s32Sum * 10;
		s32Temp = s64CheckOverflow & 0xffffffff;
		if((s32Temp / 10) != s32Sum) // overflow
		{
			return bSigned? -s32Sum : s32Sum;
		}
		else
			s32Sum = s32Sum * 10;
		if(bSigned)
		{
			if(INT_MIN + s32Sum + (*str - '0') > 0)
				return INT_MIN;
		}
		else
		{
			if(INT_MAX - s32Sum - (*str - '0') < 0)
				return INT_MAX;
		}
		s32Sum += (*str - '0');
		str++;
	}
}

int main(void)
{
	char *pstring = "-42";
	int s32Ans = myAtoi(pstring);
	printf("s32Ans = %d\n", s32Ans);
	return 1;
}
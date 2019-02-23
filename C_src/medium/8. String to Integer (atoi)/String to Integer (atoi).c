int myAtoi(char* str) {
    unsigned char bSigned = 0;
    unsigned char bCountStart = 0;
    unsigned char u8Carrynum = 0;
    int s32Sum = 0;
    long int s64CheckOverflow = 0;
    int s32Tempnum = 0;
    int s32Max = 0x7fffffff;
    int s32Min = -(0x80000000);

    while(*str)
   	{
   		if(bCountStart == 0)
   		{
   			if(*str == ' ')
   			{
   				str++;
   				continue;
   			}
   			else if(*str == '-' || *str == '+')
   			{
   				bSigned = (*str== '-') ? 1 : 0;
   				str++;
   				bCountStart = 1;
   				continue;
   			}
   			else if((*str >= '0') && (*str <= '9'))
   			{
   				s32Sum += (*str) - '0';
   				s32Sum = bSigned ? -s32Sum : s32Sum;
   				u8Carrynum++;
   				str++;
   				bCountStart = 1;
   			}
   			else // non-numerical
   			{
   				return s32Sum;
   			}
   		}
   		else // already start counting
   		{
   			if((*str >= '0') && (*str <= '9'))
   			{
   				if(u8Carrynum == 0)
   				{
	   				s32Sum += (*str) - '0';
	   				s32Sum = bSigned ? -s32Sum : s32Sum;
	   				u8Carrynum++;
	   				str++;
   				}
   				else
   				{
	   				s64CheckOverflow = (long int)s32Sum * 10;
	   				s32Tempnum  = s64CheckOverflow & 0xffffffff;
	   				if((s32Tempnum / 10) != s32Sum) // check mutiply cause overflow.
	   				{
	   					return bSigned ? s32Min : s32Max;
	   				}
	   				else
	   				{
	   					s32Sum *= 10;
	   				}
					s32Tempnum = (*str) - '0';
	   				s32Tempnum = bSigned ? -s32Tempnum : s32Tempnum;

	   				if(bSigned)
	   				{
	   					if ((s32Min - s32Sum - s32Tempnum) > 0) // out of range
	   					{
	   						return s32Min;
	   					}
	   					else
	   					{
	   						s32Sum += s32Tempnum;
	   						str++;
	   					}
	   				}
	   				else
	   				{
	   					if ((s32Max - s32Sum - s32Tempnum) < 0) // out of range
	   					{
	   						return s32Max;
	   					}
	   					else
	   					{
	   						s32Sum += s32Tempnum;
	   						str++;
	   					}
	   				}

   				}
   			}
   			else
   			{
   				return s32Sum;
   			}
   		}
   	}
   	return s32Sum;
}
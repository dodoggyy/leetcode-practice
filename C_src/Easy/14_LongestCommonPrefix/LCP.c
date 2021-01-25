/*
Runtime: 0 ms, faster than 100.00% of C online submissions for Longest Common Prefix.
Memory Usage: 7 MB, less than 100.00% of C online submissions for Longest Common Prefix.
*/

#define FALSE 0
#define TRUE 1
static unsigned char IsLocalCommonPreix(char *ComPrefix, char *InputString )
{

	if((strlen(ComPrefix) == 0) || (strlen(InputString) == 0))
	{
		ComPrefix[0] = '\0';
		return FALSE;
	}
	else
	{
		int idx = 0;
		while(ComPrefix[idx] != NULL)
		{
			if(idx == 0)
			{
				if(ComPrefix[idx] != InputString[idx])
				{
					ComPrefix[idx] = '\0';
					return FALSE; //meaning dosen't exist common prefix, break func immediately.
				}
			}
			else
			{
				if(ComPrefix[idx] != InputString[idx])
				{
					ComPrefix[idx] = 0; // update longest common prefix.
				}
			}
			idx++;
		}
	}
	return TRUE;
}

char * longestCommonPrefix(char ** strs, int strsSize){

	if((strs == NULL) || (strsSize == 0)) // input data is invaild.
	{
		return "\0";
	}


	int u32Idx = 0, u32MinStringIdx = 0;
	int u32MinStrLength = strlen(strs[u32MinStringIdx]);
	while(u32Idx < strsSize)
	{
		if(strlen(strs[u32Idx]) == 0) // input data is invaild.
		{
			return "\0";
		}
		else
		{
			u32MinStrLength = (u32MinStrLength <= strlen(strs[u32Idx])) ? u32MinStrLength : strlen(strs[u32Idx]);
			u32MinStringIdx = (u32MinStrLength <= strlen(strs[u32Idx])) ? u32MinStringIdx : u32Idx;
		}
		u32Idx++;
	}

	u32Idx = 0;
	char *ComPrefix = (char *)malloc(sizeof(int)*u32MinStrLength);
	memset(ComPrefix, 0, sizeof(int)*u32MinStrLength);
	memcpy(ComPrefix, strs[u32MinStringIdx], sizeof(char)*u32MinStrLength);

	while(u32Idx < strsSize)
	{
		if(IsLocalCommonPreix(ComPrefix, strs[u32Idx]) != TRUE)
		{
			return ComPrefix; //terminal function, cause it doesn't exist common prefix.
		}
		u32Idx ++;
	}
	return ComPrefix;
}

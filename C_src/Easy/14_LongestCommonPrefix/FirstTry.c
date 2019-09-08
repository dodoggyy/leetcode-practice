/*
Runtime: 4 ms, faster than 70.35% of C online submissions for Longest Common Prefix.
Memory Usage: 7.2 MB, less than 50.00% of C online submissions for Longest Common Prefix.
*/

char * longestCommonPrefix(char ** strs, int strsSize){

	if((strs == NULL) || (strsSize == 0)) // input data is invaild.
	{
		return "\0";
	}


	int u32Idx = 0, u32IdxString = 0;
	int u32MinStrLength = strlen(strs[0]);
	while(u32Idx < strsSize)
	{
		if(strlen(strs[u32Idx]) == 0) // input data is invaild.
		{
			return "\0";
		}
		else
		{
			u32MinStrLength = (u32MinStrLength <= strlen(strs[u32Idx])) ? u32MinStrLength : strlen(strs[u32Idx]);
		}
		u32Idx++;
	}

	u32IdxString = 0;
	u32Idx = 0;
	char *ComPrefix = (char *)malloc(sizeof(int)*u32MinStrLength);
	memset(ComPrefix, 0, sizeof(int)*u32MinStrLength);

	while(u32IdxString < u32MinStrLength)
	{
		while(u32Idx < strsSize)
		{
			if(ComPrefix[u32IdxString] == NULL)
			{

				ComPrefix[u32IdxString] = strs[u32Idx][u32IdxString];
			}
			else
			{
				if(ComPrefix[u32IdxString] == strs[u32Idx][u32IdxString])
				{
                    u32Idx ++;
					continue;
				}
				else
				{
					ComPrefix[u32IdxString] = NULL;
					return ComPrefix;
				}
			}
			u32Idx ++;
		}
		u32Idx = 0;
		u32IdxString++;
	}
	return ComPrefix;
}

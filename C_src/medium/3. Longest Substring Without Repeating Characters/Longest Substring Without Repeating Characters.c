int lengthOfLongestSubstring(char* s) {
   unsigned char u8HashTable[128] = {0}; //char size
   unsigned char u8curcount = 0;
   unsigned char u8LongestLength = 0;
   char *pu8SubString;

   while(*s)
   {
   	 pu8SubString = s;
   	 while(*pu8SubString)
   	 {
   	 	if(u8HashTable[(*pu8SubString)] != 1)
   	 	{
   	 		u8curcount++;
   	 		u8HashTable[(*pu8SubString)] = 1;
   	 		pu8SubString++;
   	 	}
   	 	else
   	 	{
   	 		if(u8LongestLength < u8curcount)
   	 		{
   	 			u8LongestLength = u8curcount;
   	 		}
   	 		u8curcount = 0;
   	 		memset(u8HashTable, 0, 128);
   	 		s++;
   	 		break;
   	 	}
   	 }
   }
   return u8LongestLength;
}
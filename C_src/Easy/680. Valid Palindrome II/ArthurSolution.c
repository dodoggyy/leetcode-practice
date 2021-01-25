bool validPalindrome(char* s) {
    int i=0;
    int StrLenth = strlen(s);
    bool LeftShift=0,RightShift=0;
    bool bRemovedChar = false;
    bool bNeedCheckRightSide = false;
    bool bDoubleCheck = false;
    do
    {
        if(s[i+LeftShift] == s[StrLenth-i-1-RightShift])
        {
            i++;
        }
        else
        {
            //Check left side first
            if(bNeedCheckRightSide == false && bRemovedChar == false && s[i+1+LeftShift] == s[StrLenth-i-1-RightShift])
            {
                LeftShift=1;
                bRemovedChar = true;
                continue;
            }
            else if(bRemovedChar == false && s[i+LeftShift] == s[StrLenth-i-2-RightShift])
            {
                RightShift=1;
                if(bNeedCheckRightSide)
                    bNeedCheckRightSide = false;
                bRemovedChar = true;
                continue;
            }
            else if(bRemovedChar && bNeedCheckRightSide == false && bDoubleCheck == false)
            {
                bDoubleCheck = true;
                LeftShift=0;
                bRemovedChar = false;
                bNeedCheckRightSide = true;
                i=i==0?0:i-1;
                continue;
            }
            else
                return false;
        }
    }while(i<StrLenth/2);
        return true;
}
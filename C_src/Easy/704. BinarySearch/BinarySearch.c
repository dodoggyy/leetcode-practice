int search(int* nums, int numsSize, int target) {
    
    if(nums == NULL || numsSize == 0)
    	return -1;
    int s32Mid;
    int s32Left = 0;
    int s32Right = numsSize;
    while(s32Left < s32Right)
  	{
  		s32Mid = (s32Left + s32Right) >> 1;
  		if(nums[s32Mid] == target)
  			return s32Mid;
  		else if(nums[s32Mid] > target)
  			s32Right = s32Mid;
  		else
  			s32Left = s32Mid + 1;
  	}
  	return -1;
}


[-1,0,3
	,5,
	9,12]
2

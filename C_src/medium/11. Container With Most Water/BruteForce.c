#define min(n,m) (n>=m)?m:n
//(Brute force)
int maxArea(int* height, int heightSize) {

	int i = 0, j = 0;
	int maxArea = 0, temp = 0;

	for(i =0; i <heightSize; i++)
		for(j=i+1; j<heightSize; j++)
		{
			temp = min(height[i],height[j]);
            temp = temp * (j-i);
			maxArea = (maxArea >= temp) ? maxArea : temp;
		}
    
    return maxArea;
}


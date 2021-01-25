//Two pointer (Greedy solution)
int maxArea(int* height, int heightSize) {

	unsigned int Left = 0, Right = heightSize - 1;
	unsigned int maxArea = 0, tempArea = 0;

	while(Left < Right)
	{
		tempArea = (height[Left] >= height[Right])? (height[Right] * (Right - Left)) : (height[Left] * (Right - Left));
		maxArea = (maxArea >= tempArea)? maxArea : tempArea;

		if(height[Left] >= height[Right])
			Right--;
		else
			Left++;
	}
	return maxArea;

}
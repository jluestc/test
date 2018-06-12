package sort1


//一、冒泡排序，每次比较相邻两个数，如果前一个数比后一个数大，则交换两个数。
//经过一次遍历后（外循环），最大数在最右侧；以此类推
func BubleSort(nums []int)[]int{
	for i:=0;i<len(nums);i++{
		for j:=1;j<len(nums)-i;j++{
			if nums[j-1]>nums[j]{
				nums[j-1],nums[j]=nums[j],nums[j-1]
			}
		}
	}
	return nums
}

//二、选择排序，对给定的数组进行多次遍历，每次均找出一个最大的或者最小值的索引。
func SelectSort(nums []int)[]int{
	length :=len(nums)
	for i:=0;i<length-1;i++{
		for j:=i+1;j<length;j++{
			if nums[j]<nums[i]{
				nums[i],nums[j]=nums[j],nums[i]
			}
		}
	}
	return nums
}
/*
func SelectSort(nums []int){
	length :=len(nums)
	for i:=0;i<lenght;i++{
		maxIndex :=0
		//寻找一个最大数，保存索引值
		for j:=1;j<length-i;j++{
			if nums[j]>nums[maxIndex]{
				maxIndex=j
			}
		}
		nums[length-i-1],nums[maxIndex]=nums[maxIndex],nums[lenght-i-1]
	}
}
*/
//三、插入排序：从第二个数开始向右遍历，每次均把该位置的元素移动至左侧，放在一个正确的位置，比左侧大，右侧小
func InsetSort(nums []int)[]int{
	for i:=1;i<len(nums);i++{
		if nums[i]<nums[i-1]{
			j :=i-1
			temp :=nums[i]
			for j>=0 && nums[j]>temp{
				nums[j+1] =nums[j]
				j--
			}
			nums[j+1]=temp
		}
	}
	return nums
}
//快排：先找到一个数pivot把数组分为两组，使得其中一组的数均大于另一组中的数//，此时pivot就是正确位置，然后对这两组数组再次进行这个操作
func QuickSort(nums []int)[]int{
	recursionSort(nums,0,len(nums)-1)
	return nums
}

func recursionSort(nums []int,left int,right int){
	if left<right{
	 	pivot :=partition(nums,left,right)
	 	recursionSort(nums,left,pivot-1)
	 	recursionSort(nums,pivot+1,right)
	}
}
func partition(nums []int,left int,right int)int{
	for left<right{
		for left<right &&nums[left]<=nums[right]{
			right--
		}
		if left<right{
			nums[left],nums[right]=nums[right],nums[left]
			left++
		}
		for left<right &&nums[left]<=nums[right]{
			left++
		}
		if left<right{
			nums[left],nums[right]=nums[right],nums[left]
			right--
		}
	}
	return left
}



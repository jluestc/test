package sort1

import(
	"testing"
	"math/rand"
	"fmt"
	
)

func TestSort1(t *testing.T){
	var tests =[][]int{{9,7,3,8,2,6,4},{3,5,4,5,6,4}}
	for _,test :=range tests{
		got1 :=BubleSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got1)
		}
/*
		got2 :=SelectSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got2)
		}

		got3 :=InsetSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got3)
		}

		got4 :=QuickSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got4)
		}*/
	}
}
func TestSort2(t *testing.T){
	var tests =[][]int{{9,7,3,8,2,6,4},{3,5,4,5,6,4}}
	for _,test :=range tests{
		
		got2 :=SelectSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got2)
		}

	}
}
func TestSort3(t *testing.T){
	var tests =[][]int{{9,7,3,8,2,6,4},{3,5,4,5,6,4}}
	for _,test :=range tests{
		got3 :=InsetSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got3)
		}
	}
}

func TestSort4(t *testing.T){
	var tests =[][]int{{9,7,3,8,2,6,4},{3,5,4,5,6,4}}
	for _,test :=range tests{	
		got4 :=QuickSort(test)
		if !Issorted(test){
			t.Errorf("nums %v want output:\n%v\n",test,got4)
		}
	}
}
//genereate input by math/rand
func TestSortQ(t *testing.T){
	var tests =[]struct{
		seed int64
		max int
		count int
	}{
		{1,100,10},
		{2,1000,10},
		{3,5000,10},
		{4,1000,30},
	}
	for _,test :=range tests{
		defer func(){
			if r :=recover(); r!=nil{
				fmt.Errorf("%v,%v",test,r)
			}
		}()

		nums :=generator(test.seed,test.max,test.count)
		got :=QuickSort(nums)
		if !Issorted(nums){
			t.Errorf("%v \nnums: %v\nsorted:%v is not ordered",test,nums,got)
		}
	}
}

func generator(seed int64,max int,count int)(res []int){
	res =make([]int,count)
	rand.Seed(seed)
	for i,_:=range res{
		res[i]=rand.Intn(max)
	}
	return
}


func Issorted(nums []int) bool{
	for i:=range nums{
		if i>0&&nums[i]<nums[i-1]{
			return false
		}
	}
	return true
}
//bench test run: go test -bench=.
func  BenchmarkSort4(b *testing.B){
	nums :=generator(1,10000,30000)
	b.ResetTimer()//?
	for i:=0;i<b.N;i++{
		QuickSort(nums)
	}
}
func Examplesort4(){
	a :=[]int{3,6,2,5,2,7,1,9,9,3,1,2,0,8}
	got :=QuickSort(a)
	if !Issorted(a){
		fmt.Printf("err: nums %v sort as:\n%v",a,got)
	}
}


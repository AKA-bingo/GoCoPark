package sort

/* 插入排序（Insertion Sort）
插入排序（Insertion-Sort）是一种简单直观的排序算法。
它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
算法描述:
插入排序不是通过交换位置而是通过比较找到合适的位置插入元素来达到排序的目的的。
1.从第一个元素开始，该元素可以认为已经被排序；
2.取出下一个元素，在已经排序的元素序列中从后向前扫描；
3.如果该元素（已排序）大于新元素，将该元素移到下一位置；
4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
5.将新元素插入到该位置后；
6.重复步骤2~5。
简单插入排序的时间复杂度也是O(n^2)。
*/
func InsertionSort(list []int64) []int64 {
	for i := 1; i < len(list); i++ { //默认第一位已经是有序，从第二位开始进行比较
		temp := list[i] //当前要进行排序的数字
		j := i - 1
		for j >= 0 && temp < list[j] { //逐一比较前面的数字，如果比当前数字大 则后移，j继续向前比较
			list[j+1] = list[j]
			j--
		}
		list[j+1] = temp //直到j停下的地方为比当前数字小的数字 在j位置的后方插入当前排序的数字
		//fmt.Println(list)
	}
	return list
}
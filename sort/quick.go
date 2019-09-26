package sort

/* 快速排序（Quick Sort）
快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，
则可分别对这两部分记录继续进行排序，以达到整个序列有序。
算法描述:
1从数列中挑出一个元素，称为 “基准”（pivot）;
2重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
*/
func QuickSort(list []int64) []int64 {
	if len(list) <= 1 {
		return list
	}

	mid, i, j := list[0], 0, len(list)-1
	for i < j {
		//	因为哨兵节点是在最左边标记 所以需要从最右边开始触发 最后交换时才能保证交点比哨兵值小
		//从右到左 找出第一个比mid小的值
		for list[j] >= mid && i < j {
			j--
		}
		list[i] = list[j] //覆盖i位置的值
		//从左到右 找出第一个比mid大的值
		for list[i] <= mid && i < j {
			i++
		}
		list[j] = list[i] //覆盖j位置的值

		//如果不覆盖 也可以用交换的方式来排序，如果i, j没相遇，交换 继续下一次
		/*if i < j {
			list[i], list[j] = list[j], list[i]
		}*/
	}

	list[i] = mid //将哨兵数字放至查找的交集点

	QuickSort(list[:i])
	QuickSort(list[i+1:])
	return list
}

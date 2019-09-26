package sort

/* 冒泡排序（Bubble Sort）
冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
算法描述:
1.比较相邻的元素。如果第一个比第二个大，就交换它们两个；
2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
3.针对所有的元素重复以上的步骤，除了最后一个；
4.重复步骤1~3，直到排序完成。
冒泡排序的时间复杂度为O(n^2)。
*/
func BubbleSort(list []int64) []int64 {
	for i := 0; i < len(list)-1; i++ { //进行len()-1次循环
		finish := true //当没有替换操作时, 表示数列有序. 退出循环
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				finish = false
			}
		}
		if finish { //替换操作标识没有改变, 表示完成
			break
		}
		//fmt.Println(list)
	}
	return list
}

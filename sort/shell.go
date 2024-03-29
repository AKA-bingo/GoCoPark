package sort

/* 希尔排序（Shell Sort）
算法描述:
先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述:
1.选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
2.按增量序列个数k，对序列进行k趟排序；
3.每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m的子序列，分别对各子表进行直接插入排序。仅增量因子为1时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/
func ShellSort(list []int64) []int64 {
	for gap := len(list) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(list); i++ {
			for j := i - gap; j >= 0 && list[j] > list[j+gap]; j -= gap {
				list[j], list[j+gap] = list[j+gap], list[j]
			}
		}
		//fmt.Println(list)
	}
	return list
}

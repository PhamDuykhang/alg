package sorting

func swap(i, j *int) {
	*i, *j = *j, *i
	return
}

func SelectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		mIdx := i
		for j := i + 1; j < len(a); j++ {
			if a[mIdx] > a[j] {
				mIdx = j
			}
		}
		swap(&a[i], &a[mIdx])
	}
}
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				swap(&a[j], &a[j+1])
			}
		}
	}
}
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		k := a[i]
		j := i
		for j > 0 && k < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = k
	}
}
func MergeSort(a []int) {
	l := len(a)
	mergeSort(a, 0, l-1)
}

func merge(a []int, l, m, r int) {
	tem := make([]int, r-l+1)

	p, q := l, m+1
	k := 0
	for i := l; i <= r; i++ {
		if p > m {
			tem[k] = a[q]
			q++
			k++
			continue
		}
		if q > r {
			tem[k] = a[p]
			p++
			k++
			continue
		}
		if a[p] < a[q] {
			tem[k] = a[p]
			p++
			k++
		} else {
			tem[k] = a[q]
			q++
			k++
		}

	}
	k=0
	for i := l; i <= r; i++ {
		a[i] = tem[k]
		k++
	}
	return

}
func mergeSort(a []int, l, r int) {

	if l < r {
		m := (r + l) / 2
		mergeSort(a, l, m)
		mergeSort(a, m+1, r)
		merge(a, l, m, r)
	}
}

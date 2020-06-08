/*
 * Copyright (c) 2020. Author by KTECH Inc team. khang_duy.p
 */

package main

import "fmt"

type (
	MaxHeep struct {
		a      []int
		height int
		length int
	}
)

func swap(i, j *int) {
	*i, *j = *j, *i
	return
}
func (m *MaxHeep) Parent(i int) int {
	return i / 2
}

func (m *MaxHeep) Left(i int) int {
	return 2 * i
}
func (m *MaxHeep) Right(i int) int {
	return 2*i + 1
}

func (m *MaxHeep) PrintArr() {
	for i := 1; i < len(m.a); i++ {
		fmt.Printf("%d ", m.a[i])
	}
}

func (m *MaxHeep) Maximum() int {
	return m.a[1]
}

func (m *MaxHeep) ExtractMax() int {
	max := m.a[1]
	m.a[1] = m.a[m.length]
	m.length--
	m.MaxHeapIfy(1)
	return max
}

func (m *MaxHeep) InsertNew(k int) error {
	m.height++
	m.a[m.height] = -9999
	err := m.IncreaseKey(m.height, k)
	if err != nil {
		return err
	}
	return nil
}

func (m *MaxHeep) IncreaseKey(i, k int) error {
	if m.a[i] > k {
		return fmt.Errorf("the number to increase must geater than old value")
	}
	m.a[i] = k
	for i > 1 && m.a[m.Parent(i)] < m.a[i] {
		swap(&m.a[m.Parent(i)], &m.a[m.Parent(i)])
		i = m.Parent(i)
	}
	return nil

}

func (m *MaxHeep) MaxHeapIfy(i int) {
	l := m.Left(i)
	r := m.Right(i)
	var largest int
	if l < m.length && m.a[l] > m.a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < m.length && m.a[r] > m.a[largest] {
		largest = r
	}
	if largest != i {
		swap(&m.a[i], &m.a[largest])
		m.MaxHeapIfy(largest)
	} else {
		return
	}
}

func (m *MaxHeep) HeepSort() {
	l := m.length - 1

	for i := l; i >= 2; i-- {
		swap(&m.a[1], &m.a[i])
		m.length = m.length - 1
		m.MaxHeapIfy(1)
	}
}
func (m *MaxHeep) BuildMaxHeep(a []int) {

	c := make([]int, len(a)+1)
	for i := 1; i < len(c); i++ {
		c[i] = a[i-1]
	}

	m.a = c
	m.length = len(a)

	for i := m.length / 2; i >= 1; i-- {
		m.MaxHeapIfy(i)
	}

}

func NewMaxHeep() *MaxHeep {
	return &MaxHeep{}
}

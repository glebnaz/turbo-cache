package list

import "testing"

func BenchmarkList_PushToFront(b *testing.B) {
	b.StopTimer()

	/////
	l := NewList(nil)
	//todo завести sync pool для того что бы аллокейтить туда element

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, err := l.PushToFront(&Element{value: i})
		if err != nil {
			panic(err)
		}
	}
}

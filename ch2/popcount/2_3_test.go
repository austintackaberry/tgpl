package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(436278463)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(436278463)
	}
}

func BenchmarkShiftCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftCount(436278463)
	}
}

func BenchmarkQuickCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickCount(436278463)
	}
}

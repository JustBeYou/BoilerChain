package main

func main() {
	store := NewInMemoryBlockStore()
	PrintBlock(store.GetTail())
}

package main

type iterator interface {
	hasNext() bool
	getNext() *User
}

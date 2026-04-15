package main

type Foo struct {
    gate_1 chan struct{};
    gate_2 chan struct{};
}

func NewFoo() *Foo {
	return &Foo{
        gate_1: make(chan struct{}),
        gate_2: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst();

    f.gate_1 <- struct{}{};
}

func (f *Foo) Second(printSecond func()) {
    <-f.gate_1;

	/// Do not change this line
	printSecond();

    f.gate_2 <- struct{}{};
}

func (f *Foo) Third(printThird func()) {
    <-f.gate_2;

	// Do not change this line
	printThird();
}

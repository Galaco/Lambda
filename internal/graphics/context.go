package graphics

import "errors"

var (
	// ErrorCallstackOutOfBounds
	ErrorCallstackOutOfBounds = errors.New("call stack out of bounds")
)

type DrawCall func()

type callStack struct {
	stack []DrawCall
}

func (ctx *callStack) Reset() {
	ctx.stack = make([]DrawCall, 0)
}

// Unwind
func (ctx *callStack) Unwind(frames int) error {
	if frames > len(ctx.stack) {
		return ErrorCallstackOutOfBounds
	}
	ctx.stack = ctx.stack[:len(ctx.stack) - frames]

	return nil
}

// Execute
func (ctx *callStack) Execute() {
	defer func() {
		if e := recover(); e != nil {
			ctx.Reset()
		}
	}()
	for _,f := range ctx.stack {
		f()
	}

	ctx.Reset()
}

// AddCall
func (ctx *callStack) AddCall(call DrawCall) {
	ctx.stack = append(ctx.stack, call)
}

// Context
type Context struct {
	Stack callStack
}
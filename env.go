package gojq

type env struct {
	pc        int
	stack     *stack
	value     []interface{}
	scopes    *stack
	codes     []*code
	codeinfos []codeinfo
	forks     []*fork
	backtrack bool
	offset    int
}

type scope struct {
	id     int
	offset int
	pc     int
}

type fork struct {
	op         opcode
	pc         int
	stackindex int
	stacklimit int
	scopeindex int
	scopelimit int
}

func newEnv() *env {
	return &env{stack: newStack(), scopes: newStack()}
}

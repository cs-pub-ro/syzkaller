// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package report

//import ("fmt")

type unikraft struct {
	*config
}

func ctorUnikraft(cfg *config) (Reporter, []string, error) {
	ctx := &unikraft{
		config: cfg,
	}
	return ctx, nil, nil
}

func (ctx *unikraft) ContainsCrash(output []byte) bool {
	//	fmt.Printf("crash\n")
	return false
}

func (ctx *unikraft) Parse(output []byte) *Report {
	//	fmt.Printf("parse\n")
	//	rep := simpleLineParser(output, unikraftOopses, unikraftStackParams, ctx.ignores)

	rep := &Report{
		Output: output,
	}

	return rep
	//	return nil
}

func (ctx *unikraft) Symbolize(rep *Report) error {
	//	fmt.Printf("symbolize\n")
	return nil
}

var unikraftOopses = []*oops{}
var unikraftStackParams = &stackParams{}

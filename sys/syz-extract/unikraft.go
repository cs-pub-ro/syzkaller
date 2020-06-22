// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/google/syzkaller/pkg/compiler"
)

type unikraft struct{}

func (*unikraft) prepare(sourcedir string, build bool, arches []*Arch) error {
	if sourcedir == "" {
		return fmt.Errorf("provide path to kernel checkout via -sourcedir flag (or make extract SOURCEDIR)")
	}
	return nil
}

func (*unikraft) prepareArch(arch *Arch) error {
	return nil
}

func (*unikraft) processFile(arch *Arch, info *compiler.ConstInfo) (map[string]uint64, map[string]bool, error) {
	params := &extractParams{
		DeclarePrintf: true,
	}

	args := []string{
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/lib/ukdebug" + "/include",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/include",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/apps/helloworld/build/include",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/arch/x86/x86_64/include",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/lib/nolibc/include",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/lib/nolibc/musl-imported/include",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/lib/nolibc/musl-imported/arch/generic",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/lib/nolibc/arch/x86_64",
		"-I" + "/home/unikraft/fuzz/unikraft/unikraft-hackathon/scripts/my_unikraft_project/unikraft/lib/uktime/include",
	}

	return extract(info, "gcc", args, params)
}

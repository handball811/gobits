package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {
	// Add Sample
	/*
		TEXT("Add", NOSPLIT, "func(x, y uint64) uint64")
		Doc("Add adds x and y.")
		x := Load(Param("x"), GP64())
		y := Load(Param("y"), GP64())
		ADDQ(x, y)
		Store(y, ReturnIndex(0))
		RET()
		Generate()
	*/

	TEXT("SlideCopy", NOSPLIT, "func(dst, src []byte, lf, rf byte, size uint64) uint64")
	Doc("Copy slided Buffer")
	dstPtr := Load(Param("dst").Base(), GP64())
	srcPtr := Load(Param("src").Base(), GP64())
	lf8 := Load(Param("lf"), GP8())
	rf8 := Load(Param("rf"), GP8())
	size := Load(Param("size"), GP64())
	ret := GP64()
	XORQ(ret, ret)

	c := GP8()
	sz := GP64()
	lf, rf := GP32(), GP32()
	XORL(lf, lf)
	XORL(rf, rf)
	MOVB(lf8, lf.As8())
	MOVB(rf8, rf.As8())

	Label("loop")

	Comment("Check If Ends")
	MOVQ(size, sz)
	SARQ(Imm(3), sz)
	CMPQ(sz, Imm(0))
	JE(LabelRef("done"))

	Comment("Copy Left Byte")

	MOVB(Mem{Base: srcPtr}, c)
	PEXTL(lf, c.As32(), c.As32())
	MOVB(c, Mem{Base: dstPtr})
	INCQ(srcPtr)

	Comment("Copy Right Byte")

	MOVB(Mem{Base: srcPtr}, c)
	PDEPL(rf, c.As32(), c.As32())
	ADDB(c, Mem{Base: dstPtr})
	INCQ(dstPtr)

	SUBQ(Imm(8), size)
	INCQ(ret)
	JMP(LabelRef("loop"))

	Label("done")

	Comment("Store copied byte length")
	Store(ret, ReturnIndex(0))

	RET()
	Generate()

}

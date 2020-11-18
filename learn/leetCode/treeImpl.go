package main

func Tree() {
	var (
		Treea, Treeb, Treec, Treed, Treee, Treef, Treeg TreeNode
	)

	Treea.Val = 7
	Treea.Left = &Treeb
	Treea.Right = &Treec

	Treeb.Val = 3
	Treeb.Left = &Treed
	Treeb.Right = &Treee

	Treec.Val = 10
	Treec.Left = &Treef
	Treec.Right = &Treeg

	Treed.Val = 1
	Treee.Val = 5
	Treef.Val = 9
	Treeg.Val = 12

}

//		   7a
//	   3b        10c
//  1d   5e  9f    12g

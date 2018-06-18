package best_travel

func ChooseBestSum(t, k int, ls []int) int {
	/*
	    V[i, w]
		  = 0            if i=0 or w=0
	      = V(i-w, w)    if wi> w
		  = V(i-w, w-wi) otherwise
	*/
}

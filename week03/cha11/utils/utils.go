package utils

// 1,1,2,3,5,8....
// 1 1 f(x-1) + (x) 
// 2 1 
// 3 2 
// 4 3 
// 5 5 
// 6 8 

func Test(n int) (res int) {
	if n == 1 || n == 2{
		res := 1
		return res
	} else {
		res := Test(n-1) +Test(n)
		return res
	}
}


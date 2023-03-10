package splitthearraytomakecoprimeproducts

func findValidSplit(nums []int) int {
	p := 1
	for i := range nums {
		p *= nums[i]
	}

	curr, res := 1, -1
    for i := 0; i < len(nums)-1;i++ {
		curr *= nums[i]
        g := gcd(curr, p/curr)
        
		if g == 1 {
			res = i
            break
		}
	}

	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

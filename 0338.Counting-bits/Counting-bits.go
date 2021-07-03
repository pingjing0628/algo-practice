package problem0338

func countBits(n int) []int {
	ans := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i >> 1] + (i & 1)
	}
	return ans
}

/* &
	0 & 0 = 0
	0 & 1 = 0
	1 & 0 = 0
	1 & 1 = 1
 */

/* Ex: 5, >> 右移, & bitwise AND
	1. i = 1		2. i = 2		3. i = 3
	   1 >> 1 = 0      2 >> 1 = 1      3 >> 1 = 1
	   >> 0001         >> 0010         >> 0011
	   -------   	   -------		   -------
	      0000			  0001			  0001
	   1 & 1 = 1       2 & 1 = 0       3 & 1 = 1
	   [0, 1]          [0, 1, 1]       [0, 1, 1, 2]
	------------    -------------    -------------
	4. i = 4		   5. i = 5
	   4 >> 1 = 2         5 >> 1 = 2
	   >> 0100            >> 0101
	   -------            -------
		  0010               0010
	   ans[2] = 1         ans[2] = 1
	   4 & 1 = 0          5 & 1 = 1
	   [0, 1, 1, 2, 1]    [0, 1, 1, 2, 1, 2]
*/

package main

func numSubmat(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				continue
			}
			if j == 0 {
				mat[i][j] = 1
			} else {
				mat[i][j] += mat[i][j-1]
			}
			ans += mat[i][j]
			t := mat[i][j]
			for k := i - 1; k >= 0; k-- {
				if mat[k][j] == 0 {
					break
				}
				if mat[k][j] < t {
					t = mat[k][j]
				}
				ans += t
			}
		}
	}
	return ans
}

func main() {

}

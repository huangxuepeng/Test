package main

import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	// 删除所有的重复的节点, 头节点也有可能被删除
// 	dummy := &ListNode{0, head}
// 	cur := dummy
// 	for cur.Next != nil && cur.Next.Next != nil {
// 		if cur.Next.Val == cur.Next.Next.Val {
// 			x := cur.Next.Val
// 			// 向后去找相等的数据
// 			for cur.Next != nil && cur.Next.Val == x {
// 				cur.Next = cur.Next.Next
// 			}
// 		} else {
// 			cur = cur.Next
// 		}
// 	}
// 	return dummy.Next
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
// 	tail := &ListNode{}
// 	carry := 0
// 	for l1 != nil || l2 != nil {
// 		n1, n2 := 0, 0
// 		if l1 != nil {
// 			n1 = l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			n2 = l2.Val
// 			l2 = l2.Next
// 		}
// 		sum := n1 + n2 + carry
// 		sum, carry = sum%10, sum/10
// 		// fmt.Println(sum)
// 		if head == nil {
// 			head = &ListNode{Val: sum}
// 			tail = head
// 		} else {
// 			tail.Next = &ListNode{Val: sum}
// 			tail = tail.Next
// 		}
// 	}
// 	if carry > 0 {
// 		tail.Next = &ListNode{Val: carry}
// 	}
// 	return head
// }

// func numSquares(n int) int {
// 	dp := make([]int, n+1)
// 	for i := range dp {
// 		dp[i] = math.MaxInt64
// 	}
// 	dp[0] = 0
// 	for i := 0; i <= n; i++ {
// 		for j := 1; j*j <= i; j++ {
// 			dp[i] = min(dp[i-j*j]+1, dp[i])
// 		}
// 	}
// 	return dp[n]
// }

// func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i-1]+nums[i] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}
// 	return max
// }
// func isSubsequence(s string, t string) bool {
// 	S, T := []byte(s), []byte(t)
// 	i, j := 0, 0
// 	for i < len(S) && j < len(T) {
// 		if S[i] == T[j] {
// 			i++
// 		}
// 		j++
// 	}
// 	return i == len(S)
// }
// func minCostClimbingStairs(cost []int) int {
// 	dp := make([]int, len(cost)+1)
// 	for i := 2; i <= len(cost); i++ {
// 		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
// 	}
// 	return dp[len(cost)]

// }
// func largestNumber(nums []int) string {
// 	sort.Slice(nums, func(i, j int) bool {
// 		x, y := nums[i], nums[j]
// 		xx, yy := 10, 10
// 		for xx <= x {
// 			xx *= 10
// 		}
// 		for yy <= y {
// 			yy *= 10
// 		}
// 		// 验证加起来之后的数据, 比较得出两个数的顺序
// 		return x*yy+y >= y*xx+x
// 	})
// 	// 验证开头为0的现象
// 	if nums[0] == 0 {
// 		return "0"
// 	}
// 	ans := []byte{}
// 	for _, v := range nums {
// 		ans = append(ans, strconv.Itoa(v)...)
// 	}
// 	return string(ans)
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func preorderTraversal(root *TreeNode) []int {
// 	var pp func(node *TreeNode)
// 	arr := []int{}
// 	pp = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		arr = append(arr, node.Val)
// 		pp(node.Left)
// 		pp(node.Right)
// 	}
// 	pp(root)
// 	return arr
// }
// func isValid(s string) bool {
// 	if len(s)%2 != 0 {
// 		return false
// 	}
// 	b := map[byte]byte{
// 		')': '(',
// 		']': '[',
// 		'}': '{',
// 	}
// 	arr := []byte{}
// 	for i := 0; i < len(s); i++ {
// 		if len(arr) == 0 {
// 			arr = append(arr, s[i])
// 		} else {
// 			if arr[len(arr)-1] == b[s[i]] {
// 				arr = arr[:len(arr)-1]
// 			} else {
// 				arr = append(arr, s[i])
// 			}
// 		}
// 	}
// 	return len(arr) == 0
// }

// // 广度优先
// func numIslands(grid [][]byte) int {
// 	res := 0
// 	m, n := len(grid), len(grid[0])
// 	if m == 0 {
// 		return 0
// 	}
// 	arr := make([][]bool, m)
// 	for i := range arr {
// 		arr[i] = make([]bool, n)
// 	}
// 	var dfs func(i, j int)
// 	dfs = func(i, j int) {
// 		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
// 			return
// 		}
// 		if arr[i][j] {
// 			return
// 		}
// 		arr[i][j] = true
// 		dfs(i-1, j)
// 		dfs(i+1, j)
// 		dfs(i, j-1)
// 		dfs(i, j+1)
// 	}

// 	for a := 0; a < m; a++ {
// 		for b := 0; b < n; b++ {
// 			if grid[a][b] == '1' && !arr[a][b] {
// 				res++
// 				dfs(a, b)
// 			}
// 		}
// 	}
// 	fmt.Println(arr)
// 	return res
// }

// func longestArithSeqLength(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	sort.Ints(nums)
// 	dp := make([][]int, len(nums))
// 	for i := range dp {
// 		dp[i] = make([]int, len(nums))
// 	}
// 	ans := 1
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < i; j++ {
// 			d := nums[i] - nums[j]
// 			dp[i][d] = max(dp[i][d], dp[j][d]+1)
// 			ans = max(ans, dp[i][d])
// 		}
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func levelOrder(root *TreeNode) [][]int {
// 	arr := [][]int{}
// 	q := []*TreeNode{root}
// 	if root == nil {
// 		return nil
// 	}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		arr = append(arr, []int{})
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			arr[i] = append(arr[i], node.Val)
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return arr
// }

// func rightSideView(root *TreeNode) []int {
// 	arr := []int{}
// 	q := []*TreeNode{root}
// 	if root == nil {
// 		return nil
// 	}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			if j == len(q)-1 {
// 				arr = append(arr, node.Val)
// 			}
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return arr
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head, tail := &ListNode{}, &ListNode{}
// 	head = tail
// 	carry := 0
// 	for l1 != nil || l2 != nil {
// 		n1, n2 := 0, 0
// 		if l1 != nil {
// 			n1 = l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			n2 = l2.Val
// 			l2 = l2.Next
// 		}
// 		sum := n1 + n2 + carry
// 		sum, carry = sum%10, sum/10
// 		tail.Next = &ListNode{Val: sum}
// 		tail = tail.Next
// 	}
// 	if carry > 0 {
// 		tail.Next = &ListNode{Val: carry}
// 	}
// 	return head.Next
// }

// pop 弹出头部的一个元素 、
// push插入一个元素
// empty 判断队列是否为空
// query 检查头部的元素
// func main() {
// 	var queen func(str string)
// 	arr := []int{}
// 	queen = func(str string) {
// 		switch str {
// 		case "push":
// 			var a int
// 			fmt.Scan(&a)
// 			arr = append(arr, a)
// 		case "pop":
// 			arr = arr[1:]
// 		case "empty":
// 			if len(arr) == 0 {
// 				fmt.Println("YES")
// 			} else {
// 				fmt.Println("NO")
// 			}
// 		case "query":
// 			fmt.Println(arr[0])
// 		}
// 	}
// 	var n int
// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		var str string
// 		fmt.Scan(&str)
// 		queen(str)
// 	}
// }

// func averageOfLevels(root *TreeNode) []float64 {
// 	if root == nil {
// 		return nil
// 	}
// 	q := []*TreeNode{root}
// 	arr := []float64{}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		sum := 0
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			sum += node.Val
// 			if j == len(q)-1 {
// 				arr = append(arr, float64(sum)/float64(len(q)))
// 			}
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return arr
// }

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

// func levelOrder(root *Node) (ans [][]int) {
// 	if root == nil {
// 		return nil
// 	}
// 	q := []*Node{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*Node{}
// 		ans = append(ans, []int{})
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			ans[i] = append(ans[i], node.Val)
// 			p = append(p, node.Children...)
// 		}
// 		q = p
// 	}
// 	return ans
// }

// func largestValues(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	arr := []int{}
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		dd := math.MinInt64
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			dd = max(dd, node.Val)
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		arr = append(arr, dd)
// 		q = p
// 	}
// 	return arr
// }

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	q := []*Node{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*Node{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			if j < len(q)-1 {
// 				node.Next = q[j+1]
// 			}
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return root
// }

// 前序遍历
// func preorderTraversal(root *TreeNode) []int {
// 	arr := []int{}
// 	var per func(node *TreeNode)
// 	per = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		arr = append(arr, node.Val)
// 		per(node.Left)
// 		per(node.Right)
// 	}
// 	per(root)
// 	return arr
// }

// 后序遍历
// func postorderTraversal(root *TreeNode) []int {
// 	arr := []int{}
// 	var per func(node *TreeNode)
// 	per = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		per(node.Left)
// 		per(node.Right)
// 		arr = append(arr, node.Val)
// 	}
// 	per(root)
// 	return arr
// }

// 中序遍历
// func inorderTraversal(root *TreeNode) []int {
// 	arr := []int{}
// 	var per func(node *TreeNode)
// 	per = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		per(node.Left)
// 		arr = append(arr, node.Val)
// 		per(node.Right)
// 	}
// 	per(root)
// 	return arr
// }

// func averageOfSubtree(root *TreeNode) int {
// 	ans := 0
// 	var dfs func(node *TreeNode) (int, int)
// 	dfs = func(node *TreeNode) (int, int) {
// 		sum, cnt := node.Val, 1
// 		if node.Left != nil {
// 			s, c := dfs(node.Left)
// 			sum += s
// 			cnt += c
// 		}
// 		if node.Right != nil {
// 			s, c := dfs(node.Right)
// 			sum += s
// 			cnt += c
// 		}
// 		if sum/cnt == node.Val {
// 			ans++
// 		}
// 		return sum, cnt
// 	}
// 	dfs(root)
// 	return ans
// }

// func largestGoodInteger(num string) string {
// 	for i := '9'; i >= '0'; i-- {
// 		s := strings.Repeat(string(i), 3)
// 		if strings.Contains(num, s) {
// 			return s
// 		}
// 	}
// 	return ""
// }

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			node.Left, node.Right = node.Right, node.Left
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return root
// }

// func lengthOfLongestSubstring(s string) int {
// 	res := 0
// 	for i := 0; i < len(s); i++ {
// 		arr := map[byte]bool{}
// 		ans := 0
// 		for j := i; j < len(s); j++ {
// 			if arr[s[i]] {
// 				break
// 			} else {
// 				arr[s[i]] = true
// 				ans++
// 			}
// 		}
// 		res = max(res, ans)
// 	}
// 	return res
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type Node struct {
// 	Val      int
// 	Children []*Node
// }

// func preorder(root *Node) []int {
// 	var per func(node *Node)
// 	arr := []int{}
// 	per = func(node *Node) {
// 		if node == nil {
// 			return
// 		}
// 		arr = append(arr, node.Val)
// 		for _, v := range node.Children {
// 			arr = append(arr, v.Val)
// 		}
// 	}
// 	per(root)
// 	return arr
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
// func isSymmetric(root *TreeNode) bool {
// 	var pp func(left, right *TreeNode) bool
// 	pp = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		} else if left != nil && right == nil ||
// 			left == nil && right != nil || left.Val != right.Val {
// 			return false
// 		}
// 		return pp(left.Left, right.Right) && pp(left.Right, right.Left)
// 	}
// 	if root == nil {
// 		return true
// 	}
// 	return pp(root.Left, root.Right)
// }

// func removeDuplicates(s string) string {
// 	arr := []byte{}
// 	for i := 0; i < len(s); i++ {
// 		if len(arr) > 0 && s[i] == arr[len(arr)-1] {
// 			arr = arr[:len(arr)-1]
// 		} else {
// 			arr = append(arr, s[i])
// 		}
// 	}
// 	return string(arr)
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		var a, b int
// 		fmt.Scan(&a, &b)
// 		if b-a >= 2 {
// 			ans++
// 		}
// 	}
// 	fmt.Println(ans)
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func maxDepth(root *TreeNode) int {
// 	var per func(node *TreeNode) int
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	per = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left, right := per(node.Left), per(node.Right)
// 		return 1 + max(left, right)
// 	}
// 	return per(root)
// }

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	deep := 0
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		deep++
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return deep
// }

// type Node struct {
// 	Val      int
// 	Children []*Node
// }

// func maxDepth(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}
// 	deep := 0
// 	q := []*Node{root}
// 	for i := 0; len(q) > 0; i++ {
// 		deep++
// 		p := []*Node{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			if node.Children != nil {
// 				p = append(p, node.Children...)
// 			}
// 		}
// 		q = p
// 	}
// 	return deep
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func minDepth(root *TreeNode) int {
// 	min := func(a, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}
// 	var minRoot func(node *TreeNode) int
// 	minRoot = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left, right := minRoot(node.Left), minRoot(node.Right)
// 		if node.Left == nil && node.Right != nil {
// 			return 1 + right
// 		}
// 		if node.Right == nil && node.Left != nil {
// 			return 1 + left
// 		}
// 		return 1 + min(left, right)
// 	}
// 	if root == nil {
// 		return 0
// 	}
// 	return minRoot(root)
// }
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := 0
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			res++
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}

// 		}
// 		q = p
// 	}
// 	return res
// }

// func countNodes(root *TreeNode) int {
// 	var per func(node *TreeNode) int
// 	per = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left, right := per(node.Left), per(node.Right)
// 		return 1 + left + right
// 	}
// 	if root == nil {
// 		return 0
// 	}
// 	return per(root)
// }

// func maxDepth(root *TreeNode) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	var per func(node *TreeNode) int
// 	per = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left, right := per(node.Left), per(node.Right)
// 		return max(left, right) + 1
// 	}
// 	return per(root)

// }

// func main() {
// 	go a()
// 	go b()
// 	ch := make(chan int)
// 	ch <- 1
// }

// func a() {
// 	for {
// 		fmt.Println("here1")
// 		time.Sleep(time.Second * 1)
// 	}
// }
// func b() {
// 	for {
// 		fmt.Println("here2")
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func main() {
// 	// a := "aaa"
// 	// ssh := *(*reflect.SliceHeader)(unsafe.Pointer(&a))
// 	// b := *(*[]byte)(unsafe.Pointer(&ssh))
// 	// fmt.Println(b)
// 	var a sync.Map
// 	a.Store("address", map[string]string{"pro": "江苏", "city": "南京"})
// 	v, _ := a.Load("address")
// 	fmt.Println(v.(map[string]string)["city"])
// }

// type User struct{}
// type User1 User
// type User2 = User

// func (i User) m1() {
// 	fmt.Println("m1")
// }
// func (i User1) m1() {
// 	fmt.Println("m1")
// }
// func (i User) m2() {
// 	fmt.Println("m2")
// }
// func main() {
// 	var i1 User1
// 	// var i2 User2
// 	i1.m1()
// 	// i2.m2()
// 	// i2.m1()
// }
// func main() {
// 	ch := make(chan int, 100)
// 	// A
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 	}()
// 	// B
// 	go func() {
// 		for {
// 			a, ok := <-ch
// 			if !ok {
// 				fmt.Println("close")
// 				return
// 			}
// 			fmt.Println("a: ", a)
// 		}
// 	}()
// 	// close(ch)
// 	fmt.Println("ok")
// 	time.Sleep(time.Second * 10)
// }

// select 处理异步操作, 每一个case语句必须是io操作   单向channel不允许关闭

/*
	用户注册: 首次操作需要进行的操作, 需要用户提供自己手机号码, 邮箱, 昵称, 头像等基本信息, 通过前端出道后端
	使用http协议的 post请求, 后端拿到数据之后对数据进行存储数据库, 对用户的密码进行md5加密为了防止彩虹表找出密码,
	我们可以对用户的密码进行加盐而后md5加密进行存入数据库.
	用户登录: 用户根据注册的信息, 登录网站, 后端会向前端发送token, 前端除了要使用token做一个令牌之外, 还可以生成
	一个cookie, 用于验证是否是用户本人的操作, 后端使用中间件, 验证上一个url是否属于本网站, 防止xxs注入.
	实名注册: 当学生或者企业需要拿出证明自己身份的证明, 并且传回照片, 后端调用第三方接口, 对图片信息进行识别,
	并且自动审核输入的信息和上传图片信息是否一致, 不一致, 直接发邮箱告知用户, 如果通过了, 管理员进行人工审核,
	判断信息是否真正符合要求, 并且, 管理员的id会存入被审核的用户信息里, 让每一步操作都可查
	信息审核: 用户的实名信息, 以及企业的招聘信息, b端都可以查询, 并且审核是否符合规范, 如果不符合规范就发邮箱
	提示操作失败, 需要重新发布,
*/

/*
	I think animals should be used for experimentation,
because for the benefit of humans, animal experimentation cannot be prohibited.
	Without them, a lot of advances in the field of medicine would not have been possible.
Animals are ideal subjects for experimentation.
	The first reason is that animals are biologically similar to humans,
and they are also susceptible to diseases similar to humans.
 The second is that because they have short life-cycles,
 scientists can study their entire life spans or even generations.
 The third reason is that scientists are able to perform experiments on animals that control variables,
 such as controlling their diet, temperature, lighting, etc.,
 which is difficult to achieve if people are studied.
Although animal experiments can also cause animal suffering,
 scientists have also proposed some ways to reduce animal suffering.
Although it is unethical to use animals for experiments,
the benefits behind this are still very objective for humans.
*/

// type ConfigOne struct {
// 	Daemon string
// }

// func (d ConfigOne) Strings() string {
// 	return fmt.Sprintf("print:sss %v", d)
// }
// func main() {
// 	c := ConfigOne{}
// 	ss := c.Strings()
// 	fmt.Print(ss)
// }

// func main() {
// 	var a = []int{1, 2, 3, 4, 5}
// 	var r = make([]int, 0)
// 	for i, v := range a { // 这个遍历的只是a的副本
// 		if i == 0 {
// 			a = append(a, 6, 7)
// 		}
// 		r = append(r, v)
// 	}
// 	fmt.Println(r)
// 	fmt.Println(a)
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// var path []string

// func binaryTreePaths(root *TreeNode) []string {
// 	if root == nil {
// 		return nil
// 	}
// 	pp(root, "")
// 	return path
// }
// func pp(root *TreeNode, str string) {

// 	if root != nil {
// 		paths := str
// 		paths += strconv.Itoa(root.Val)
// 		if root.Left == nil && root.Right == nil {
// 			path = append(path, paths)
// 		} else {
// 			paths += "->"
// 			pp(root.Left, paths)
// 			pp(root.Right, paths)
// 		}
// 	}
// }
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func isSymmetric(root *TreeNode) bool {
// 	var pp func(left, right *TreeNode) bool
// 	pp = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		} else if left != nil && right == nil || left == nil && right != nil || left.Val != right.Val {
// 			return false
// 		}
// 		return pp(left.Left, right.Right) && pp(left.Right, right.Left)
// 	}
// 	if root == nil {
// 		return true
// 	}
// 	return pp(root.Left, root.Right)
// }
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	var pp func(left, right *TreeNode) bool
// 	pp = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		} else if left != nil && right == nil || left == nil && right != nil || left.Val != right.Val {
// 			return false
// 		}
// 		return pp(left.Left, right.Left) && pp(left.Right, right.Right)
// 	}
// 	return pp(p, q)
// }
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
// 	var pp func(left, right *TreeNode) bool
// 	pp = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		} else if left != nil && right == nil || left == nil && right != nil || left.Val != right.Val {
// 			return false
// 		}
// 		return pp(left.Left, right.Left) && pp(left.Right, right.Right)
// 	}
// 	if root == nil {
// 		return false
// 	}
// 	return pp(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
// }
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func maxDepth(root *TreeNode) int {
// 	// 层序解决
// 	res := 0
// 	if root == nil {
// 		return 0
// 	}
// 	q := []*TreeNode{root}
// 	for i := 0; len(q) > 0; i++ {
// 		p := []*TreeNode{}
// 		res++
// 		for j := 0; j < len(q); j++ {
// 			node := q[j]
// 			if node.Left != nil {
// 				p = append(p, node.Left)
// 			}
// 			if node.Right != nil {
// 				p = append(p, node.Right)
// 			}
// 		}
// 		q = p
// 	}
// 	return res
// }
// func maxDepth(root *TreeNode) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	var pp func(node *TreeNode) int
// 	pp = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left, right := pp(node.Left), pp(node.Right)
// 		return max(left, right) + 1
// 	}
// 	if root == nil {
// 		return 0
// 	}
// 	return pp(root)
// }
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func minDepth(root *TreeNode) int {
// 	min := func(a, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}
// 	var pp func(node *TreeNode) int
// 	pp = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left, right := pp(node.Left), pp(node.Right)
// 		if node.Left != nil && node.Right == nil {
// 			return 1 + left
// 		}
// 		if node.Left == nil && node.Right != nil {
// 			return 1 + right
// 		}
// 		return 1 + min(left, right)
// 	}
// 	if root == nil {
// 		return 0
// 	}
// 	return pp(root)
// }

// func longestPalindrome(s string) string {
// 	ss := func(str string, left, right int) (int, int) {
// 		for ; left >= 0 && right < len(s) && str[left] == str[right]; left, right = left-1, right+1 {
// 		}
// 		return left + 1, right - 1
// 	}
// 	if s == "" {
// 		return ""
// 	}
// 	left, right := 0, 0
// 	for i := 0; i < len(s); i++ {

// 		left1, right1 := ss(s, i, i)
// 		left2, right2 := ss(s, i, i+1)
// 		if right1-left1 > right-left {
// 			left, right = left1, right1
// 		}
// 		if right2-left2 > right-left {
// 			left, right = left2, right2
// 		}
// 	}
// 	return s[left : right+1]
// }
// 33. 搜索旋转排序数组
// func search(nums []int, target int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return -1
// 	}
// 	if n == 1 {
// 		if target == nums[0] {
// 			return 0
// 		}
// 		return -1
// 	}
// 	l, r := 0, n-1
// 	for r >= l {
// 		mid := (l + r) >> 1
// 		if nums[mid] == target {
// 			return mid
// 		}
// 		if nums[0] <= nums[mid] { // 前半段是有序的
// 			if nums[mid] > target && nums[0] <= target {
// 				r = mid - 1
// 			} else {
// 				l = mid + 1
// 			}
// 		} else {
// 			if target > nums[mid] && target <= nums[n-1] {
// 				l = mid + 1
// 			} else {
// 				r = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }
// 最长的递增子数组
// func lengthOfLIS(nums []int) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	dp := make([]int, len(nums))
// 	maa := 1
// 	for i := range dp {
// 		dp[i] = 1
// 	}
// 	for i := 1; i < len(nums); i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		maa = max(dp[i], maa)
// 	}
// 	return maa
// }

// func mySqrt(x int) int {
// 	if x == 0 || x == 1 {
// 		return x
// 	}
// 	l, r, ans := 1, x, -1
// 	for r >= l {
// 		mid := (l + r) >> 1
// 		if mid*mid == x {
// 			return mid
// 		} else if mid*mid > x {
// 			r = mid - 1
// 		} else {
// 			ans = mid
// 			l = mid + 1
// 		}
// 	}
// 	return ans
// }

// func search(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	if len(nums) == 1 {
// 		if nums[0] == target {
// 			return 0
// 		}
// 		return -1
// 	}
// 	l, r := 0, len(nums)-1
// 	for r >= l {
// 		mid := (l + r) >> 1
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] > target {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }

// func findLength(nums1 []int, nums2 []int) int {

// }
// func main() {
// 	x := 1
// 	fmt.Println(x)
// 	{
// 		fmt.Println(x)
// 		i, x := 2, 2
// 		fmt.Println(i, x)
// 	}
// 	fmt.Println(x) // print ?
// }

// type People struct {
// 	Name string `json:"name"`
// }

// func main() {
// 	js := `{
// 		"name":"seekload",
// 	}`
// 	var p People
// 	err := json.Unmarshal([]byte(js), &p)
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}
// 	fmt.Println(p)
// }
// func main() {
// 	fmt.Println("kk")
// }
// func main() {
// 	isMatch := func(i int) bool {
// 		switch i {

// 		case 1, 2:
// 			return true
// 		}
// 		return false
// 	}
// 	fmt.Println(isMatch(1))
// 	fmt.Println(isMatch(2))
// }
// func main() {
// 	// var fn1 = func(a, b int) {
// 	// 	if a > b {
// 	// 		fmt.Println(a)
// 	// 	}
// 	// }
// 	var ff func()
// 	if ff == nil {
// 		fmt.Println("iii")
// 	} else {
// 		fmt.Println("kk")
// 	}
// }
type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	fmt.Println(n)
	n++
	n.test()
	// fmt.Println(f1)
	n++
}

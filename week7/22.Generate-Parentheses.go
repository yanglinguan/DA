func generateParenthesis(n int) []string {
    res := []string{}
    helper(n, 0, "", &res)
    return res
}

func helper(left int, right int, cur string, res *[]string) {
    if left == 0 && right == 0 {
        *res = append(*res, cur)
        return
    }

    if left > 0 {
        helper(left-1, right+1, cur + "(", res)
    }

    if right > 0 {
        helper(left, right-1, cur + ")", res)
    }
}

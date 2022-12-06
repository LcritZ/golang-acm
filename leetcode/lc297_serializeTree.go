package leetcode

import (
    "golang-acm/util"
    "strconv"
    "strings"
)

/**
二叉树按层遍历，输出字符串，并且可以恢复成节点

其实当做数组做，找到下标的规律就可以恢复
 */

type Codec struct{}

func ConstructorTree() (_ Codec) {
    return
}

func (Codec) serialize(root *util.TreeNode) string {
    sb := &strings.Builder{}
    var dfs func(*util.TreeNode)
    dfs = func(node *util.TreeNode) {
        if node == nil {
            sb.WriteString("null,")
            return
        }
        sb.WriteString(strconv.Itoa(node.Val))
        sb.WriteByte(',')
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return sb.String()
}

func (Codec) deserialize(data string) *util.TreeNode {
    sp := strings.Split(data, ",")
    var build func() *util.TreeNode
    build = func() *util.TreeNode {
        if sp[0] == "null" {
            sp = sp[1:]
            return nil
        }
        val, _ := strconv.Atoi(sp[0])
        sp = sp[1:]
        return &util.TreeNode{val, build(), build()}
    }
    return build()
}

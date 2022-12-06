package leetcode

import "golang-acm/util"

func Flatten(root *util.TreeNode)  {
   if root == nil {
       return
   }
   Flatten(root.Left)
   Flatten(root.Right)

   l := root.Left
   r := root.Right

   root.Left = nil
   root.Right = l
   node := root
   for node.Right != nil {
       node = node.Right
   }

   node.Right = r


}

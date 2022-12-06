package basic

import "golang-acm/util"

//BinTree* treeToList(BinTree *root)
//{
//stack<BinTree*> s;
//BinTree *p = root, *prev = NULL, *head = NULL;
//while (p || !s.empty()) {
//while (p) {
//s.push(p);
//p = p->left;
//}
//
//p = s.top();
//s.pop();
//if(head == NULL) head = p;
//
//if (prev) prev->right = p;
//p->left = prev;
//prev = p;
//
//p = p->right;
//}
//return head;
//}

//
//class Solution:
//def Convert(self, root):
//if not root:
//return None
//if not root.left and not root.right:
//return root
//
//# 将左子树构建成双链表，返回链表头
//left = self.Convert(root.left)
//p = left
//
//# 定位至左子树的最右的一个结点
//while left and p.right:
//p = p.right
//
//# 如果左子树不为空，将当前root加到左子树链表
//if left:
//p.right = root
//root.left = p
//
//# 将右子树构造成双链表，返回链表头
//right = self.Convert(root.right)
//# 如果右子树不为空，将该链表追加到root结点之后
//if right:
//right.left = root
//root.right = right
//
//return left if left else root

/**
     5
   3    7
 1   4 6  8

*/
func Tree2List(root *util.TreeNode) *util.TreeNode {
    if root == nil {
        return root
    }

    stack := []*util.TreeNode{}
    var head, prev *util.TreeNode
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        if head == nil {
            head = node
        }
        if prev != nil {
            prev.Right = node
        }
        node.Left = prev
        prev = node
        node = node.Right
        stack = stack[:len(stack)-1]
    }

    return head
}



package leetcode

/**
输入: sx = 1, sy = 1, tx = 3, ty = 5
输出: true
解释:
可以通过以下一系列转换从起点转换到终点：
(1, 1) -> (1, 2)
(1, 2) -> (3, 2)
(3, 2) -> (3, 5)

通过逆推的方法，直到其中一个相等，然后判断另一个值差值是否可以等于 这个相等的值的倍数
 */
func ReachingPoints(sx int, sy int, tx int, ty int) bool {
    if tx == sx && ty == sy {
        return true
    }
    if tx == ty {
        return false
    }

    for tx > sx && ty > sy {
        if tx < ty {
            ty = ty-tx
        } else {
            tx = tx-ty
        }
    }
    if tx < sx || ty < sy {
        return false
    } else if tx == sx {
        if ty == sy {
            return true
        }
        temp := ty-sy
        if temp%tx == 0 {
            return true
        } else {
            return false
        }
    } else if ty == sy {
        if tx == sx {
            return true
        }
        temp := tx-sx
        if temp%ty == 0 {
            return true
        } else {
            return false
        }
    }

    return false
}

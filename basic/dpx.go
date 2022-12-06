package basic

/**
5个节点

每个节点对应任务，节点之间有依赖

串行执行，满足依赖

a,b
a,c
b,d
c,d
e
     b
a         d
     c
e

a  b/c  d

e


*/

//func selectOpt(dep [][]byte) []byte {
//    res := []byte{}
//    if len(dep) == 0 {
//        return res
//    }
//
//    resMap := make(map[byte]bool)
//    nodeVal := make(map[byte]int)
//
//    for i := 0; i < len(dep); i++ {
//        for j := 0; j < len(dep[i]); j++ {
//
//
//        }
//
//    }
//
//
//}




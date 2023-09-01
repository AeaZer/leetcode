package breadthfirstsearch

// 433. 最小基因变化
// https://leetcode.cn/problems/minimum-genetic-mutation/

/*基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。*/

// 思路：假设 bank 中的元素都是合法的
func minMutation(start, end string, bank []string) int {
    if start == end {
        return 0
    }
    bankSet := map[string]struct{}{}
    for _, s := range bank {
        bankSet[s] = struct{}{}
    }
    if _, ok := bankSet[end]; !ok {
        return -1
    }

    q := []string{start}
    for step := 0; q != nil; step++ {
        tmp := q
        q = nil
        for _, cur := range tmp {
            for i, x := range cur {
                for _, y := range "ACGT" {
                    if y != x {
                        nxt := cur[:i] + string(y) + cur[i+1:]
                        if _, ok := bankSet[nxt]; ok {
                            if nxt == end {
                                return step + 1
                            }
                            delete(bankSet, nxt)
                            q = append(q, nxt)
                        }
                    }
                }
            }
        }
    }
    return -1
}

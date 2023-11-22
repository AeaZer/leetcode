package huawei

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 题目大概是输入有两行，第一行是 cities ip 的范围配置使用半角逗号隔开，第二行是给出的 ips，输出 ips 属于最佳的城市
// cities =  "city1=1.1.1.1,1.1.1.2;city2=1.1.1.1,1.1.1.4"
// ipsStr = "1.1.1.1,1.1.1.3"
// 输出：city1,city2
// 输出解释：ip=1.1.1.1 city1 和 city2 都符合但是 city1 范围更小所以返回 city1 而不是 city2
type city struct {
	name       string
	start, end int64
	distance   int64
}

func PerfectCities(cities string, ipsStr string) {
	pairs := strings.Split(cities, ";")
	citiesConfig := make([]*city, 0, len(pairs))
	for i := range pairs {
		str := strings.SplitN(pairs[i], "=", 2)
		rg := strings.SplitN(str[1], ",", 2)
		start, end := ipStrToInt(rg[0]), ipStrToInt(rg[1])
		citiesConfig = append(citiesConfig, &city{
			name:     str[0],
			start:    start,
			end:      end,
			distance: end - start,
		})
	}
	sort.Slice(citiesConfig, func(i, j int) bool {
		return citiesConfig[i].start < citiesConfig[j].start
	})
	ips := strings.Split(ipsStr, ",")
	for i := 0; i < len(ips); i++ {
		name := perfectCity(ips[i], citiesConfig)
		if i == 0 {
			fmt.Print(name)
		} else {
			fmt.Printf(",%s", name)
		}
	}
}

func ipStrToInt(str string) int64 {
	arr := strings.SplitN(str, ".", 4)
	var ans int64
	for _, s := range arr {
		i, _ := strconv.ParseInt(s, 10, 64)
		ans <<= 8
		ans |= i
	}
	return ans
}

func perfectCity(ip string, citiesConfig []*city) string {
	ipInt := ipStrToInt(ip)
	var cityName string
	var minDis int64 = math.MaxInt
	for i, _ := range citiesConfig {
		if citiesConfig[i].start > ipInt {
			break
		}
		if citiesConfig[i].start <= ipInt && citiesConfig[i].end >= ipInt && minDis > citiesConfig[i].distance {
			cityName = citiesConfig[i].name
			minDis = citiesConfig[i].distance
		}
	}
	return cityName
}

package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs = []string{"", ""}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	flag := true // that word not in var ret
	for _, v := range strs {
		flag = true
		for kRst, vRst := range ret {
			vRst_tmp := vRst[0]
			if len(vRst_tmp) == len(v) {

				// the empty value of letter
				if len(vRst_tmp) == 0 {
					ret[kRst] = append(ret[kRst], v)
					flag = false
					break
				}
				for k := range v {
					for kk := range vRst_tmp {
						if vRst_tmp[kk:kk+1] == v[k:k+1] {
							//redefined var vRst_tmp :remove the letter has been used
							if kk == 0 {
								vRst_tmp = vRst_tmp[1:]
							} else {
								if kk == len(v) {
									vRst_tmp = vRst_tmp[0 : len(v)-2]
								} else {
									vRst_tmp = vRst_tmp[0:kk] + vRst_tmp[kk+1:]
								}
							}
							break
						}
					}
					if vRst_tmp == "" {
						flag = false
						ret[kRst] = append(ret[kRst], v)
					}
				}
			}
		}
		if flag {
			ret = append(ret, []string{v})
		}
	}

	return ret
}

---
title: K-th Smallest Prime Fraction
date: 2020-01-19 12:09:01
categories: [算法,leetcode]
tags: [算法,leetcode]
description: 算法,leetcode
---

### 来源
https://leetcode.com/problems/k-th-smallest-prime-fraction/

### 描述
在包含1的数组中，两两相除（p/q且p<q）生成一个列表，获取列表第k小的两个数字（p,q）

Examples:
Input: A = [1, 2, 3, 5], K = 3
Output: [2, 5]
Explanation:
1/5, 1/3, 2/5, 1/2, 3/5, 2/3.
第三小的是 2/5.

Input: A = [1, 7], K = 1
Output: [1, 7]

<!--more-->

#### 算法
```php
// 时间复杂度超过限制，max(n,k)nlogn
class kthSplHeap extends SplHeap
{
    public function compare($a, $b) // 最小堆
    {
        return (Solution::$A[$b[0]] * Solution::$A[$a[1]]) - (Solution::$A[$a[0]] * Solution::$A[$b[1]]);
    }
}

class Solution
{

    private $lenA = 0;
    static $A = [];

    /**
     * @param Integer[] $A
     * @param Integer   $K
     * @return Integer[]
     */
    function kthSmallestPrimeFraction($A, $K)
    {
        $this->lenA = count($A);
        self::$A    = $A;
        $kthSplHeap = new kthSplHeap();
        for ($key = 0; $key < $this->lenA - 1; $key++) {
            $kthSplHeap->insert([$key, $this->lenA - 1]);
        }

        for ($k = 0; $k < $K - 1; $k++) {
            $top = $kthSplHeap->extract();
            if ($top[1] - 1 > $top[0]) {
                $kthSplHeap->insert([$top[0], $top[1] - 1]);
            }
        }

        $top = $kthSplHeap->top();

        return [$A[$top[0]], $A[$top[1]]];
    }
}

$t = time();
var_dump((new Solution())->kthSmallestPrimeFraction([],1144765));
var_dump(time() - $t);
// 1/5 1/3 2/5 1/2 2/3 3/5
die();

var_dump((new Solution())->kthSmallestPrimeFraction([1, 13, 17, 59], 6));
// 1/59 1/17 1/13 13/59 17/59 13/17
die();

```
```php
/**
二分法，找出中间点
*/
class Solution
{

    /**
     * @param Integer[] $A
     * @param Integer   $K
     * @return Integer[]
     */
    function kthSmallestPrimeFraction($A, $K)
    {
        $start = 0;
        $end   = 1;
        $len   = count($A);
        $p     = $q = 0;
        while ($start < $end) {
            $mid     = ($start + $end) / 2;
            $maxRate = 0;
            $count   = 0;

            $can_add = 0;
            for ($i = 0; $i < $len - 1; $i++) {
                while ($can_add < $len and $A[$i] / $A[$can_add] > $mid) {
                    $can_add++;
                }
                if ($can_add == $len) {
                    break;
                }
                $count += $len - $can_add;
                if ($A[$i] / $A[$can_add] > $maxRate) {
                    $maxRate = $A[$i] / $A[$can_add];
                    $p       = $A[$i];
                    $q       = $A[$can_add];
                }
            }

            if ($count == $K) {
                return [$p, $q];
            } elseif ($count > $K) {
                $end = $mid;
            } else {
                $start = $mid;
            }
        }

    }
}


```



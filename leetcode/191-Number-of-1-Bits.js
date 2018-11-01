/**
 * 191. Number of 1 Bits
 * Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

 Example 1:

 Input: 11
 Output: 3
 Explanation: Integer 11 has binary representation 00000000000000000000000000001011
 Example 2:

 Input: 128
 Output: 1
 Explanation: Integer 128 has binary representation 00000000000000000000000010000000
 */
/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function (n) {
    if (n === 0) {
        return 0;
    }
    num = 0;
    while (true) {
        let maxBin_ = maxBin(n);
        if (maxBin_.left === 0) {
            num++;
            return num;
        } else {
            n = n - maxBin_.max;
            num++;
        }
    }
};

function maxBin(m) {
    let max = 1;
    let left = m;
    while (true) {
        if (max > 1 && max >= m) {
            if (max > m) {
                max = max / 2;
            }
            left = m - max;
            return {max: max, left: left}
        } else if (max === 1 && max === m) {
            max = 1;
            left = 0;
            return {max: max, left: left}
        } else {
            max = max * 2;
        }
    }
}


var log = function (s) {
    console.log(s)
};
log(hammingWeight(99999999));



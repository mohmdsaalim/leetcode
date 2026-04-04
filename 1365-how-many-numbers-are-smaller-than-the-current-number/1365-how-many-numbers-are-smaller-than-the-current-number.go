// var smallerNumbersThanCurrent = function(nums) {
//     let sorted = [...nums].sort((a, b) => a - b);
//     let map = new Map();


//     for (let i = 0; i < sorted.length; i++) {
//         if (!map.has(sorted[i])) {
//             map.set(sorted[i], i);
//         }
//     }

//     // Build answer
//     return nums.map(num => map.get(num));
// };



func smallerNumbersThanCurrent(nums []int) []int {
    const K = 101
    freq := make([]int, K)
    // Count frequencies
    for _, x := range nums {
        freq[x]++
    }
    // Build prefix sum: prefix[v] = # of elements < v
    prefix := make([]int, K)
    for v := 1; v < K; v++ {
        prefix[v] = prefix[v-1] + freq[v-1]
    }
    // Answer for each num is prefix[num]
    res := make([]int, len(nums))
    for i, x := range nums {
        res[i] = prefix[x]
    }
    return res
}
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(a, b) {
    const map = new Map();
    for (let i = 0; i < a.length; i++) {
        let diff = b - a[i];
        if (map.has(diff)) {
            return [map.get(diff), i];
        }
        map.set(a[i], i);
    }
};
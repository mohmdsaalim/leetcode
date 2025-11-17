/**
 * @param {string} allowed
 * @param {string[]} words
 * @return {number}
 */
var countConsistentStrings = function(allowed, words) {
    return   words.filter(w => [...w].every(c => allowed.includes(c))).length;
};
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    const a = s.trim();
    const words = a.split(/\s+/);
    const lastWord = words[words.length - 1];
    return lastWord.length;
};
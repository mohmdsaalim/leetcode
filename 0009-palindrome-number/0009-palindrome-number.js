/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
   let str = x+"";
   return str === str.split("").reverse().join("")

}


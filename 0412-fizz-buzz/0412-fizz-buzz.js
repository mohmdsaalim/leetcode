/**
 * @param {number} n
 * @return {string[]}
 */
var fizzBuzz = function(n) {
    let result = [];

    for (let a = 1; a <= n; a++) {
        if (a % 3 === 0 && a % 5 === 0) {
            result.push("FizzBuzz");
        } else if (a % 3 === 0) {
            result.push("Fizz");
        } else if (a % 5 === 0) {
            result.push("Buzz");
        } else {
            result.push(`${a}`);
        }
    }

   return result;

};

console.log(fizzBuzz(15));
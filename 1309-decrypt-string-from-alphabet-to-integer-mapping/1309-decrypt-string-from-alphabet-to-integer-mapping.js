/**
 * @param {string} s
 * @return {string}
 */
var freqAlphabets = s =>
s.replace(/\d{2}#|\d/g, m => String.fromCharCode(96 + parseInt(m)));
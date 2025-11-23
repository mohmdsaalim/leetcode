var findWords = function(words) {
    const row1 = new Set("qwertyuiop");
    const row2 = new Set("asdfghjkl");
    const row3 = new Set("zxcvbnm");

    return words.filter(word => {
        const w = new Set(word.toLowerCase());
        return (
            [...w].every(ch => row1.has(ch)) ||
            [...w].every(ch => row2.has(ch)) ||
            [...w].every(ch => row3.has(ch))
        );
    });
};
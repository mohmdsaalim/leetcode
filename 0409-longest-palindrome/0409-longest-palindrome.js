var longestPalindrome = function(s) {
    const map = new Map();
    
    // Count frequency of each character
    for (const ch of s) {
        map.set(ch, (map.get(ch) || 0) + 1);
    }

    let length = 0;
    let oddFound = false;

    // Calculate palindrome length
    for (const count of map.values()) {
        if (count % 2 === 0) {
            length += count;      // use full even counts
        } else {
            length += count - 1;  // use even part of odd count
            oddFound = true;
        }
    }

    // If odd char available, we can place exactly 1 in the center
    if (oddFound) length += 1;

    return length;
};
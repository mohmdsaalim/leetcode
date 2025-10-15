var myAtoi = function(str) {
    let i = 0, sign = 1, num = 0;
    const MAX = 2147483647, MIN = -2147483648;


    while (i < str.length && str[i] === ' ') i++;


    if (str[i] === '-' || str[i] === '+') {
        sign = str[i] === '-' ? -1 : 1;
        i++;
    }


    while (i < str.length && str[i] >= '0' && str[i] <= '9') {
        num = num * 10 + (str[i].charCodeAt(0) - 48);

        if (sign * num > MAX) return MAX;
        if (sign * num < MIN) return MIN;
        i++;
    }

    return sign * num;
};
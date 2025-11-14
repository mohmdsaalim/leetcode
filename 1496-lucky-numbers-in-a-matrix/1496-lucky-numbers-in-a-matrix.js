/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var luckyNumbers  = function(matrix) {
    const m = matrix.length;
    const n = matrix[0].length;
    const result = [];

    for (let i = 0; i < m; i++) {
      
        let rowMin = matrix[i][0];
        let colIndex = 0;

        for (let j = 1; j < n; j++) {
            if (matrix[i][j] < rowMin) {
                rowMin = matrix[i][j];
                colIndex = j;
            }
        }

        let isMaxInCol = true;
        for (let k = 0; k < m; k++) {
            if (matrix[k][colIndex] > rowMin) {
                isMaxInCol = false;
                break;
            }
        }

        if (isMaxInCol) result.push(rowMin);
    }

    return result;
};
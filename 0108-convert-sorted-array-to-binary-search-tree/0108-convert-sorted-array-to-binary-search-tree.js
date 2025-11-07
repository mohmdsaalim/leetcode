/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
var sortedArrayToBST = function(nums) {
  // Helper recursive function
  function buildBST(left, right) {
    // Base case — if range invalid, return null
    if (left > right) return null;

    // Pick middle element as root (to keep it balanced)
    const mid = Math.floor((left + right) / 2);
    const root = new TreeNode(nums[mid]);

    // Recursively build left and right subtrees
    root.left = buildBST(left, mid - 1);
    root.right = buildBST(mid + 1, right);

    return root;
  }

  return buildBST(0, nums.length - 1);
};
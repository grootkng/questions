/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function (nums, target) {
  let head = 0

  for (head; head <= nums.length; head++) {
    if (nums[head] === target || target < nums[head]) {
      break
    } else if (nums[head] < target && nums[head+1] > target) {
      head = head + 1
      break
    } else if (head == nums.length-1 && nums[head] !== target) {
      head = nums.length
      break
    }
  }

  return head
};

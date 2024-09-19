# 485. Max Consecutive Ones

[Leetcode #485](https://leetcode.com/problems/max-consecutive-ones/description/)

## Description

Given a binary array `nums`, return the maximum number of consecutive `1`'s in the array.

Example 1:
```
Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
```

Example 2:
```
Input: nums = [1,0,1,1,0,1]
Output: 2
```

Constraints:
- `1` <= `nums.length` <= `10**5`
- `nums[i]` is either `0` or `1`.

## Solution

To solve this problem, we'll use a simple linear traversal approach. The idea is to iterate through the array and keep track of the current number of consecutive ones and update the maximum count whenever we find a longer sequence.

**Approach:**

1. **Initialize two counters:**
   - `currentCount` to keep track of the number of consecutive ones encountered so far.
   - `maxCount` to store the maximum number of consecutive ones found during the traversal.

2. **Iterate through the array:**
   - If the current element is `1`, increment `currentCount` by `1`.
   - If `currentCount` exceeds `maxCount`, update `maxCount` with the value of `currentCount`.
   - If the current element is `0`, reset `currentCount` to `0` because the sequence of consecutive ones has been broken.

3. **Return `maxCount` after completing the traversal.**

This approach ensures we only go through the array once, achieving an O(n) time complexity, where n is the length of the array.

Here's the Go implementation of the solution:

```go
func findMaxConsecutiveOnes(nums []int) int {
    maxCount := 0
    currentCount := 0
    for _, num := range nums {
        if num == 1 {
            currentCount++
            if currentCount > maxCount {
                maxCount = currentCount
            }
        } else {
            currentCount = 0
        }
    }
    return maxCount
}
```

**Explanation:**

- **Variables Initialization:**
  - `maxCount` is initialized to `0` to store the maximum number of consecutive ones found.
  - `currentCount` is initialized to `0` to count the current streak of consecutive ones.

- **Iteration:**
  - We use a `for` loop to iterate over each element `num` in the `nums` array.
  - If `num` is `1`, we increment `currentCount`. We then check if `currentCount` is greater than `maxCount`, and if so, we update `maxCount`.
  - If `num` is not `1` (i.e., it's `0`), we reset `currentCount` to `0` because the consecutive streak has been broken.

- **Result:**
  - After the loop completes, `maxCount` holds the maximum number of consecutive ones found in the array, which we return as the result.

**Test Cases:**

1. **Example 1:**

   - **Input:** `[1,1,0,1,1,1]`
   - **Output:** `3`
   - **Explanation:** The last three digits are consecutive ones. The function correctly returns `3`.

2. **Example 2:**

   - **Input:** `[1,0,1,1,0,1]`
   - **Output:** `2`
   - **Explanation:** The maximum number of consecutive ones is `2`. The function correctly returns `2`.

**Time and Space Complexity:**

- **Time Complexity:** O(n), where n is the length of the array, because we traverse the array only once.
- **Space Complexity:** O(1), since we use a constant amount of extra space.
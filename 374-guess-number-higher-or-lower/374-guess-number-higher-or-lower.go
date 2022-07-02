/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		guess := guess(mid)
		if guess == -1 {
			right = mid - 1
		} else if guess == 1 {
			left = mid + 1
		} else if guess == 0 {
			return mid
		}
	}
	return left
}
import "sort"

func fourSum(nums []int, target int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }

    sort.Ints(nums)
    ans := [][]int{}
    length := len(nums)

    for i := 0; i < length; i++ {
        // avoid accessing duplicate values
        if i > 0 && nums[i - 1] == nums[i] {
            continue
        }

        for j := i + 1; j < length; j++ {
            if j > i + 1 && nums[j - 1] == nums[j] {
                continue
            }
            left := j + 1
            right := length - 1
            preSum := nums[i] + nums[j]

            for left < right {
                sum := preSum + nums[left] + nums[right]
                if sum < target {
                    left++
                } else if sum > target {
                    right--
                } else {
                    ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
                    for left + 1 < length && nums[left] == nums[left + 1] {
                        left++
                    }
                    left++
                    for right - 1 >= 0 && nums[right] == nums[right - 1] {
                        right--
                    }
                    right--
                }
            }
        }
    }

    return ans
}
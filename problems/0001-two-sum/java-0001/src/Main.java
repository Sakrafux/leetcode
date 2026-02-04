import java.util.Arrays;
import java.util.HashMap;

public class Main {

    public static void main(String[] args) {
        var solution = new Solution();

        System.out.println(Arrays.toString(solution.twoSum(new int[]{2, 7, 11, 15}, 9)));
        System.out.println(Arrays.toString(solution.twoSum(new int[]{3, 2, 4}, 6)));
        System.out.println(Arrays.toString(solution.twoSum(new int[]{3, 3}, 6)));
    }

}

class Solution {

    public int[] twoSum(int[] nums, int target) {
        var map = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
        }

        for (int i = 0; i < nums.length; i++) {
            var diff = map.get(target - nums[i]);
            if (diff != null && diff != i) {
                return new int[]{i, diff};
            }
        }

        return new int[0];
    }

}
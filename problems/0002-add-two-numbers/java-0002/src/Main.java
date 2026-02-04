public class Main {

    public static void main(String[] args) {
        runCase(new int[]{2, 4, 3}, new int[]{5, 6, 4});
        runCase(new int[]{0}, new int[]{0});
        runCase(new int[]{9,9,9,9,9,9,9}, new int[]{9,9,9,9});
    }

    private static void runCase(int[] a1, int[] a2) {
        var l1 = new ListNode();
        var l2 = new ListNode();

        var l1Cursor = l1;
        for (int j : a1) {
            l1Cursor.next = new ListNode(j);
            l1Cursor = l1Cursor.next;
        }
        l1 = l1.next;

        var l2Cursor = l2;
        for (int j : a2) {
            l2Cursor.next = new ListNode(j);
            l2Cursor = l2Cursor.next;
        }
        l2 = l2.next;

        var output = new Solution().addTwoNumbers(l1, l2);
        System.out.println(output);
    }

}

class Solution {

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        var next1 = l1;
        var next2 = l2;

        var result = new ListNode();
        var nextResult = result;

        while (next1 != null && next2 != null) {
            var tmp = next1.val + next2.val;
            nextResult.next = new ListNode(tmp % 10);
            nextResult = nextResult.next;
            next1 = next1.next;
            next2 = next2.next;

            if (tmp / 10 == 1) {
                if (next1 != null) {
                    next1.val = next1.val + 1;
                } else if (next2 != null) {
                    next2.val = next2.val + 1;
                } else {
                    nextResult.next = new ListNode(1);
                }
            }
        }

        while (next1 != null) {
            var val = next1.val;

            nextResult.next = new ListNode(val % 10);
            nextResult = nextResult.next;
            next1 = next1.next;

            if (val / 10 == 1) {
                if (next1 != null) {
                    next1.val = next1.val + 1;
                } else {
                    nextResult.next = new ListNode(1);
                }
            }
        }

        while (next2 != null) {
            var val = next2.val;

            nextResult.next = new ListNode(val % 10);
            nextResult = nextResult.next;
            next2 = next2.next;

            if (val / 10 == 1) {
                if (next2 != null) {
                    next2.val = next2.val + 1;
                } else {
                    nextResult.next = new ListNode(1);
                }
            }
        }

        return result.next;
    }

}

class ListNode {

    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    public String toString() {
        if (next == null) {
            return Integer.toString(val);
        }
        return val + "," + next;
    }

}
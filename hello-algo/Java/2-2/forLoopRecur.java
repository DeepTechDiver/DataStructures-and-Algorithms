package hello.algo;


import java.util.Stack;

public class forLoopRecur {
    /* 使用迭代模拟递归 */
    int forLoopRecur(int n) {
        // 使用一个显式的栈来模拟系统调用栈
        Stack<Integer> stack = new Stack<>();
        int res = 0;
        // 递：递归调用
        for (int i = n; i > 0; i--) {
            // 通过“入栈操作”模拟“递”
            stack.push(i);
        }
        // 归：返回结果
        while (!stack.isEmpty()) {
            // 通过“出栈操作”模拟“归”
            res += stack.pop();
        }
        // res = 1+2+3+...+n
        return res;
    }

    public static void main(String[] args) {
        forLoopRecur f = new forLoopRecur();
        System.out.println("1+2+3+...+n = " + f.forLoopRecur(5));
    }
    
}
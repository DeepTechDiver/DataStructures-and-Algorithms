package org.example.algorithm.recursion_single;

/**
 * 递归反向打印字符串
 */
public class E02ReversePrintString {

    public static void reversePrint(String str, int index) {
        if (index == str.length()) {
            return;
        }
        reversePrint(str, index + 1);
        System.out.print(str.charAt(index));
    }

    public static void main(String[] args)
    {
        reversePrint("abcdefg",0);
    }
}



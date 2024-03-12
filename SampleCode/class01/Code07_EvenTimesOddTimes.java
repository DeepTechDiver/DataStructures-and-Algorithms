package class01;

public class Code07_EvenTimesOddTimes {


    // 求数组中一种数为奇数次，其余为偶数次
    public static void printOddTimesNum(int[] arr) {
        int eor = 0;
        for (int cur : arr) {
            eor = eor ^ cur;
        }
        System.out.println(eor);

    }

    // 求数组中两种数为奇数次，其余为偶数次
    public static void printOddTimesNum2(int[] arr) {
        int eor = 0;
        for (int curNum : arr) {
            eor = eor ^ curNum;
        }
        /*
            eor  =  a ^ b
            eor != 0
            eor 必然有一个位置上是1
        */
        int onlyOne = 0; // eor'
        int rightOne = eor & (~eor + 1); // 提取出最右的1
        for (int cur : arr){
            if ((cur & rightOne) == 1 ){ // 找到 那个位 上为0（1）的数，即两边取一边
                onlyOne ^= cur; // onlyOne=a（b），(eor ^ onlyOne)=b（a）
            }
        }
        System.out.println(onlyOne + " " + (eor ^ onlyOne));
    }



    public static void main(String[] args) {
        printOddTimesNum(new int[]{2,1,1,1,3,3,3,2,1});
        printOddTimesNum2(new int[]{2,1,1,1,3,3,3,2,1,4,4,4});
    }
}

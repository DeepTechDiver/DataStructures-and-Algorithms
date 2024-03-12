package class01;


import java.util.Arrays;

public class Code02_BubbleSort {

    public static void bubbleSort(int[] arr) {
        if (arr == null || arr.length < 2){
            return;
        }
        for (int e = arr.length -1; e > 0; e--){
            for (int i = 0; i < e; i++){
                if (arr[i] > arr[i + 1]){
                    swap(arr, i, i + 1);
                }
            }
        }

    }

//    如果指向同一个即操作了同一个元素，指向了同一个内存地址，这种方法不能实现交换。 只限于操作俩个不同的内存地址
    public static void swap(int[] arr, int i,int j) {
        arr[i] = arr[i] ^ arr[j];
        arr[j] = arr[i] ^ arr[j];
        arr[i] = arr[i] ^ arr[j];
    }


    public static void main(String[] args) {
        int[] arr = {5,4,3,2,1};

        // 调用bubbleSort函数对数组进行排序
        bubbleSort(arr);

        int[] arr2 = {5,4,3,2,1};
        swap(arr2, 0,0);

        // 输出交换后的数组
        System.out.println(Arrays.toString(arr));
        System.out.println(Arrays.toString(arr2));

//        int a = 1;
//        int b = 2;
//        a ^= b ^= a ^= b;
//        System.out.println(a + "---" + b);
    }
}



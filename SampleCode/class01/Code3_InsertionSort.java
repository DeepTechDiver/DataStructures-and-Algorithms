package class01;

import java.util.Arrays;

public class Code3_InsertionSort {
    public static void insertionSort(int[] arr) {
        if (arr == null || arr.length < 2){
            return;
        }
        // 0~0 有序, 实现0~i 有序
        for (int i = 1; i < arr.length; i++) {
            for (int j = i -1; j >= 0 && arr[j] > arr[j+1]; j--){
                swap(arr, j, j+1);
            }
        }
    }
    public static void swap(int[] arr, int i,int j) {
        arr[i] = arr[i] ^ arr[j];
        arr[j] = arr[i] ^ arr[j];
        arr[i] = arr[i] ^ arr[j];
    }

    public static void main(String[] args) {
        int[] arr = {3,2,5,4,2,3,3};
        insertionSort(arr);
        System.out.println(Arrays.toString(arr));
    }

}

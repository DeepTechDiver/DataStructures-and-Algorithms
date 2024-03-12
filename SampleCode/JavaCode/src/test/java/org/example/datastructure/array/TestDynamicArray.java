package org.example.datastructure.array;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.awt.*;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.function.Consumer;

import static org.junit.jupiter.api.Assertions.*;

public class TestDynamicArray {
    @Test
    @DisplayName("测试添加")
    public void test() {
        DynamicArray dynamicArray = new DynamicArray();
        dynamicArray.addLast(1);
        dynamicArray.addLast(2);
        dynamicArray.addLast(3);
        dynamicArray.addLast(4);

        dynamicArray.add(1, 7);

        for (int i = 0; i < 8; i++) {
            System.out.println(dynamicArray.get(i));
        }
    }

    @Test
    @DisplayName("测试遍历1")
    public void testForeach() {
        DynamicArray dynamicArray = new DynamicArray();
        dynamicArray.addLast(1);
        dynamicArray.addLast(2);
        dynamicArray.addLast(3);
        dynamicArray.addLast(4);
        dynamicArray.addLast(5);
        dynamicArray.add(1, 7);

        dynamicArray.foreach((element) -> {
            System.out.println(element); // element 接收 array[i]
        });
    }

    @Test
    @DisplayName("测试遍历2")
    public void testIterator() {
        DynamicArray dynamicArray = new DynamicArray();
        dynamicArray.addLast(1);
        dynamicArray.addLast(2);
        dynamicArray.addLast(3);
        dynamicArray.addLast(4);
        dynamicArray.addLast(5);
        dynamicArray.add(1, 7);
        // 迭代器
        for (Integer element : dynamicArray) { // hashNext() next()
            System.out.println(element);
        }
    }

    @Test
    @DisplayName("测试遍历3")
    public void testStream() {
        DynamicArray dynamicArray = new DynamicArray();
        dynamicArray.addLast(1);
        dynamicArray.addLast(2);
        dynamicArray.addLast(3);
        dynamicArray.addLast(4);
        dynamicArray.addLast(5);
        dynamicArray.add(1, 7);
        // 迭代器
        dynamicArray.stream().forEach((element) -> {
            System.out.println(element);
        });
//        assertArrayEquals(
//                new int[]{1,2,3,4,5},
//                dynamicArray.stream().toArray()
//        );
    }

    @Test
    @DisplayName("测试删除")
    public void testRemove() {
        DynamicArray dynamicArray = new DynamicArray();
        dynamicArray.addLast(1);
        dynamicArray.addLast(2);
        dynamicArray.addLast(3);
        dynamicArray.addLast(4);
        dynamicArray.addLast(5);

        int removed = dynamicArray.remove(2);

        System.out.println(removed);
        System.out.println("------------------");
        dynamicArray.stream().forEach((element) -> {
            System.out.println(element);
        });
        assertEquals(3, removed);
        assertIterableEquals(Arrays.asList(1, 2, 3, 4), dynamicArray);
    }

    @Test
    @DisplayName("测试扩容")
    public void testAddLast() {
        DynamicArray dynamicArray = new DynamicArray();
        for (int i = 0; i < 9; i++) {
            dynamicArray.addLast(i + 1);
        }
        assertIterableEquals(
                Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9),
                dynamicArray
        );
    }


}

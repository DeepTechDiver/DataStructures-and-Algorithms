package org.example.datastructure.linkedlist;


import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;
import static org.junit.jupiter.api.Assertions.*;


public class TestSinglyLinkedList {

    private SinglyLinkedList getLinkedList() {
        SinglyLinkedList list = new SinglyLinkedList();
        list.addLast(0);
        list.addLast(1);
        list.addLast(2);
        list.addLast(3);
        list.addLast(4);
        return list;
    }

    @Test
    @DisplayName("测试 addFirst")
    public void test1() {
        SinglyLinkedList list = new SinglyLinkedList();
        list.addFirst(1);
        list.addFirst(2);
        list.addFirst(3);
        list.addFirst(4);
        list.loop1(i -> System.out.println(i));
    }
    @Test
    @DisplayName("测试")
    public void test2() {
        SinglyLinkedList list = new SinglyLinkedList();
        list.addFirst(0);
        list.addFirst(1);
        list.addFirst(2);
        list.addLast(3);
        list.loop1(i -> System.out.println(i));
        assertIterableEquals(Arrays.asList(0, 1, 2, 3),list);
    }

    @Test
    @DisplayName("测试 get")
    public void test3() {
        SinglyLinkedList list = getLinkedList();
        assertEquals(2, list.get(2));
        assertThrows(IllegalArgumentException.class, () -> list.get(10));
    }


    @Test
    @DisplayName("测试 removeFirst")
    public void test4() {
        SinglyLinkedList list = getLinkedList();

        list.removeFirst();
        assertIterableEquals(Arrays.asList(1, 2, 3, 4), list);
        assertThrows(IllegalArgumentException.class, list::removeFirst);
    }


    @Test
    @DisplayName("测试 remove")
    public void test5() {
        SinglyLinkedList list = getLinkedList();

        list.remove(3);
        assertIterableEquals(Arrays.asList(0, 1, 2, 4), list);
//        assertThrows(IllegalArgumentException.class, list::removeFirst);
    }
}

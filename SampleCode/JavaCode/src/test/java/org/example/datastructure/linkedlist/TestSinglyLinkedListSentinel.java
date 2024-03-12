package org.example.datastructure.linkedlist;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class TestSinglyLinkedListSentinel {
    private SinglyLinkedListSentinel getLinkedList() {
        SinglyLinkedListSentinel list = new SinglyLinkedListSentinel();
        list.addLast(0);
        list.addLast(1);
        list.addLast(2);
        list.addLast(3);
        list.addLast(4);
        return list;
    }


    @Test
    @DisplayName("测试 get")
    public void test3() {
        SinglyLinkedListSentinel list = getLinkedList();
        assertEquals(2, list.get(2));
        assertThrows(IllegalArgumentException.class, () -> list.get(10));
    }

}

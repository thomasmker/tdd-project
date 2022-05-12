package com.thomashayashi;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class TestMoney {
    @Test
    public void testMultiplication() {
        Dollar five= new Dollar(5);
        five.times(2);
        assertEquals(10, five.amount);
    }

    class Dollar {
        int amount;
        public Dollar(int amount) {
            this.amount = amount;
        }

        void times(int multipler) {
            amount *= multipler;
        }
    }
}

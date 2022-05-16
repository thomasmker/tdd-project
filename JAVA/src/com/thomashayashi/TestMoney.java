package com.thomashayashi;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public class TestMoney {
    @Test
    public void testMultiplication() {
        Money five= new Dollar(5, "USD");
        assertEquals(new Dollar(10, "USD"), five.times(2));
        assertEquals(new Dollar(15, "USD"), five.times(3));
    }

    @Test
    public void testFrancMultiplication() {
        Money five= new Franc(5, "CHF");
        assertEquals(new Franc(10, "CHF"), five.times(2));
        assertEquals(new Franc(15, "CHF"), five.times(3));
    }

    @Test
    public void testEquality() {
        assertTrue(new Dollar(5, "USD").equals(new Dollar(5, "USD")));
        assertFalse(new Dollar(5, "USD").equals(new Dollar(6, "USD")));
        assertTrue(new Franc(5, "CHF").equals(new Franc(5, "CHF")));
        assertFalse(new Franc(5, "CHF").equals(new Franc(6, "CHF")));
        assertFalse(new Franc(5, "CHF").equals(new Dollar(5, "USD")));
    }

    @Test
    public void testCurrency() {
        assertEquals("USD", new Dollar(1, "USD").currency());
        assertEquals("CHF", new Franc(1, "CHF").currency());
    }

    @Test
    public void testDifferentClassEquality() {
        assertTrue(new Money(10, "CHF").equals(new Franc(10, "CHF")));
    }
}


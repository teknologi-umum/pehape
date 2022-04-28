package com.teknologiumum.pehape;

import static org.junit.Assert.assertArrayEquals;

import org.junit.Before;
import org.junit.Test;
import com.teknologiumum.pehape.String.PehapeString;

/**
 * Unit test for simple App.
 */
public class AppTest {
    /**
     * PehapeString.explode should explode like PHP explode()
     */
    @Test
    public void shouldExplodeLikePhp() {
        assertArrayEquals(new String[] { "Hello mama" }, PehapeString.explode(" ", "Hello mama", 0));
        assertArrayEquals(new String[] { "Hello" }, PehapeString.explode(" ", "Hello mama", -1));
        assertArrayEquals(new String[] { "Hello", "mama" }, PehapeString.explode(" ", "Hello mama", 2));
        assertArrayEquals(new String[] {}, PehapeString.explode(" ", "Hello mama", -3));
        assertArrayEquals(new String[] {}, PehapeString.explode(" ", "Hello mama", -2));
        assertArrayEquals(new String[] { "Hello" }, PehapeString.explode(" ", "Hello mama", -1));
    }
}

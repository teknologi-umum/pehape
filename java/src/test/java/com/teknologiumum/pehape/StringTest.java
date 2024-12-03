package com.teknologiumum.pehape;

import static org.junit.Assert.assertArrayEquals;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;
import org.junit.rules.ExpectedException;

import com.teknologiumum.pehape.String.PehapeString;

/**
 * Unit test for simple App.
 */
public class StringTest {
    /**
     * PehapeString.explode should explode like PHP explode()
     */
    @Test
    public void shouldExplodeLikePhpSingleSpace() {
        String content = "Hello pehape world";
        assertArrayEquals(new String[] { "Hello", "pehape", "world" }, PehapeString.explode(" ", content));
        assertArrayEquals(new String[] { "Hello pehape world" }, PehapeString.explode(" ", content, 0));

        assertArrayEquals(new String[] { "Hello pehape world" }, PehapeString.explode(" ", content, 1));
        assertArrayEquals(new String[] { "Hello", "pehape world" }, PehapeString.explode(" ", content, 2));
        assertArrayEquals(new String[] { "Hello", "pehape", "world" }, PehapeString.explode(" ", content, 3));

        assertArrayEquals(new String[] { "Hello", "pehape" }, PehapeString.explode(" ", content, -1));
        assertArrayEquals(new String[] { "Hello" }, PehapeString.explode(" ", content, -2));
        assertArrayEquals(new String[] {}, PehapeString.explode(" ", content, -3));

        assertArrayEquals(new String[] { "Hello pehape world" }, PehapeString.explode(" ", content, null));
    }

    @Test
    public void shouldExplodeLikePhpDoubleSpace() {
        String content = "Hello pehape world";
        assertArrayEquals(new String[] { "Hello pehape world" }, PehapeString.explode("  ", content));

    }

    @Test()
    public void shouldExplodeLikePhpWithException() {
        String content = "Hello pehape world";

        Assert.assertThrows(IllegalArgumentException.class, () -> {
            PehapeString.explode(null, content); // should throw IllegalArgument
        });

        Assert.assertThrows(IllegalArgumentException.class, () -> {
            PehapeString.explode(" ", null); // should throw IllegalArgument
        });

        Assert.assertThrows(IllegalArgumentException.class, () -> {
            PehapeString.explode("", content); // should throw IllegalArgument
        });

        Assert.assertThrows(IllegalArgumentException.class, () -> {
            PehapeString.explode(" ", null, 1); // should throw IllegalArgument
        });

        Assert.assertThrows(IllegalArgumentException.class, () -> {
            PehapeString.explode(null, content, 1); // should throw IllegalArgument
        });

        Assert.assertThrows(IllegalArgumentException.class, () -> {
            PehapeString.explode("", content, 1); // should throw IllegalArgument
        });

    }
}

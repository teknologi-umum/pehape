package org.teknologiumum.pehape;

import static org.junit.Assert.assertArrayEquals;
import static org.junit.Assert.assertTrue;

import java.lang.System.Logger;
import java.util.logging.LogManager;

import org.junit.Before;
import org.junit.Test;
import org.teknologiumum.pehape.String.PehapeString;

/**
 * Unit test for simple App.
 */
public class AppTest 
{
    private String content;
    private String separator;
    @Before
    public void initValue() {
        content = "Hello mama";
        separator = " ";
    }
    /**
     * Rigorous Test :-)
     */
    @Test
    public void shouldAnswerWithTrue()
    {
        assertArrayEquals(new String[]{"Hello mama"}, PehapeString.explode(separator, content, 0));
        assertArrayEquals(new String[]{"Hello"}, PehapeString.explode(" ", "Hello mama", -1));
        assertArrayEquals(new String[]{"Hello", "mama"}, PehapeString.explode(" ", "Hello mama", 2));
        assertArrayEquals(new String[]{}, PehapeString.explode(" ", "Hello mama", -3));
        assertArrayEquals(new String[]{}, PehapeString.explode(" ", "Hello mama", -2));
        assertArrayEquals(new String[]{"Hello"}, PehapeString.explode(" ", "Hello mama", -1));
    }
}

package com.teknologiumum.pehape.String;

import java.util.Arrays;

public class PehapeString {

    public static String[] explode(String separator, String content) {

        if (separator.length() == 0 || separator == null) {
            throw new IllegalArgumentException("separator can't be empty");
        }

        if (content.length() == 0 || content == null) {
            throw new IllegalArgumentException("content can't be empty");
        }

        return content.split(separator);
    }

    public static String[] explode(String separator, String content, Integer limit) {

        if (limit == null) {
            return explode(content, separator);
        }

        if (separator.length() == 0 || separator == null) {
            throw new IllegalArgumentException("separator can't be empty");
        }

        if (content.length() == 0 || content == null) {
            throw new IllegalArgumentException("content can't be empty");
        }

        if (limit > 0) {
            return content.split(separator, limit);
        } else if (limit < 0) {
            String[] array = content.split(separator, limit);
            if (Math.abs(limit) > array.length) {
                return Arrays.copyOfRange(array, 0, 0);
            } else {
                return Arrays.copyOfRange(array, 0, array.length + limit);
            }

        } else {
            return content.split(separator, 1);
        }
    }
}

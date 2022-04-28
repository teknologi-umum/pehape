package com.teknologiumum.pehape.String;

import java.util.Arrays;

import org.apache.commons.lang3.StringUtils;

public class PehapeString {

    public static String[] explode(String separator, String content) {

        if (StringUtils.isEmpty(separator)) {
            throw new IllegalArgumentException("Separator can't be empty.");
        }

        if (StringUtils.isEmpty(content)) {
            throw new IllegalArgumentException("Content can't be empty.");
        }

        return content.split(separator);
    }

    public static String[] explode(String separator, String content, Integer limit) {

        if (limit == null) {
            return explode(separator, content);
        }

        if (StringUtils.isEmpty(separator)) {
            throw new IllegalArgumentException("Separator can't be empty.");
        }

        if (StringUtils.isEmpty(content)) {
            throw new IllegalArgumentException("Content can't be empty.");
        }

        if (limit > 0) {
            return content.split(separator, limit);
        } else if (limit < 0) {
            String[] array = content.split(separator, limit);
            if (Math.abs(limit) >= array.length) {
                return new String[0];
            } else {
                return Arrays.copyOfRange(array, 0, array.length + limit);
            }
        } else {
            return content.split(separator, 1);
        }
    }
}

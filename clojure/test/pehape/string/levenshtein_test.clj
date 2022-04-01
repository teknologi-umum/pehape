(ns pehape.string.levenshtein-test
  (:require [clojure.test :refer [deftest is testing]]
            [pehape.core :refer [levenshtein]]))


(deftest levenshtein-test
  (testing "Throws an error when str1 or str2 is nil"
    (is (thrown-with-msg? IllegalArgumentException #"str1 cannot be nil"
                          (levenshtein nil "")))
    (is (thrown-with-msg? IllegalArgumentException #"str1 cannot be nil"
                          (levenshtein nil "" 1 1 1)))
    (is (thrown-with-msg? IllegalArgumentException #"str2 cannot be nil"
                          (levenshtein "" nil)))
    (is (thrown-with-msg? IllegalArgumentException #"str2 cannot be nil"
                          (levenshtein "" nil 1 1 1))))
  (testing "Throws an error when cost is negative"
    (is (thrown-with-msg? IllegalArgumentException #"insert-cost cannot be a negative value"
                          (levenshtein "a" "a" -1 1 1)))
    (is (thrown-with-msg? IllegalArgumentException #"replace-cost cannot be a negative value"
                          (levenshtein "a" "a" 1 -1 1)))
    (is (thrown-with-msg? IllegalArgumentException #"delete-cost cannot be a negative value"
                          (levenshtein "a" "a" 1 1 -1))))
  (testing "Returns Correct Distance"
    (testing "Custom Test Cases"
      ;; equal
      (is (= 0 (levenshtein "12345" "12345")))
      ;; first string empty
      (is (= 3 (levenshtein "" "xyz")))
      ;; second string empty
      (is (= 3 (levenshtein "xyz" "")))
      (is (= 3 (levenshtein "xyz" "" 1 1 1)))
      ;; both empty
      (is (= 0 (levenshtein "" "")))
      (is (= 0 (levenshtein "" "" 10 10 10)))
      ;; 1 character
      (is (= 1 (levenshtein "1" "2")))
      ;; 2 characters swapped
      (is (= 2 (levenshtein "12" "21")))
      ;; inexpensive deletion
      (is (= 2 (levenshtein "2121" "11" 2 1 1)))
      ;; expensive deletion
      (is (= 10 (levenshtein "2121" "11" 2 1 5)))
      ;; inexpensive insertion
      (is (= 2 (levenshtein "11" "2121")))
      ;; expensive insertion
      (is (= 10 (levenshtein "11" "2121" 5 1 1)))
      ;; inexpensive replacement
      (is (= 3 (levenshtein "111" "121" 2 3 2)))
      ;; very expensive replacement
      (is (= 4 (levenshtein "111" "121" 2 9 2))))
    ;; test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_error_conditions.phpt
    (testing "PHP Test Cases"
      ;; levenshtein no longer has a maximum string length limit
      (is (= 254 (levenshtein "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsu" "A")))
      (is (= 255 (levenshtein "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuv" "A")))
      (is (= 254 (levenshtein "A" "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsu")))
      (is (= 255 (levenshtein "A" "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuv"))))
    ;; test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_bug_16473.phpt
    (testing "PHP Bug #16473 Test Cases"
      (is (= 2 (levenshtein "a" "bc")))
      (is (= 2 (levenshtein "xa" "xbc")))
      (is (= 2 (levenshtein "xax" "xbcx")))
      (is (= 2 (levenshtein "ax" "bcx"))))
    ;; test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_bug_6562.phpt
    (testing "PHP Bug #6562 Test Cases"
      (is (= 1 (levenshtein "debug" "debugg")))
      (is (= 1 (levenshtein "ddebug" "debug")))
      (is (= 2 (levenshtein "debbbug" "debug")))
      (is (= 1 (levenshtein "debugging" "debuging"))))
    ;; test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_bug_7368.phpt
    (testing "PHP Bug #7368 Test Cases"
      (is (= 2 (levenshtein "13458" "12345")))
      (is (= 2 (levenshtein "1345" "1234"))))))
(ns pehape.string-test
  (:require [clojure.test :refer [deftest is testing]]
            [pehape.string :refer [explode implode]]))


(deftest explode-test
  (testing "Can Explode Space Separated String By Space"
    #_{:clj-kondo/ignore [:inline-def]}
    (def content "Hello pehape world")
    (is (= (explode " " content) ["Hello" "pehape" "world"]))
    (is (= (explode " " content 0) ["Hello"]))
    (is (= (explode " " content 1) ["Hello"]))
    (is (= (explode " " content 5) ["Hello" "pehape" "world"]))
    (is (= (explode " " content -1) ["Hello" "pehape"]))
    (is (= (explode " " content -3) [])))
  (testing "Can Explode Double Space Separated String By Space"
    #_{:clj-kondo/ignore [:inline-def]}
    (def content "Hello  pehape  world")
    (is (= (explode " " content) ["Hello" "" "pehape" "" "world"])))
  (testing "Can Explode Space Separated String By Double Space"
    #_{:clj-kondo/ignore [:inline-def]}
    (def content "Hello pehape world")
    (is (= (explode "  " content) ["Hello pehape world"])))
  (testing "Explode Space Separated String By Empty String Or Nil Throws Exception"
    #_{:clj-kondo/ignore [:inline-def]}
    (def content "Hello pehape world")
    (is (thrown-with-msg? IllegalArgumentException #"separator cannot be nil or empty" (explode "" content)))
    (is (thrown-with-msg? IllegalArgumentException #"separator cannot be nil or empty" (explode nil content)))
    (is (thrown-with-msg? IllegalArgumentException #"separator cannot be nil or empty" (explode "" content 1)))
    (is (thrown-with-msg? IllegalArgumentException #"separator cannot be nil or empty" (explode nil content 1))))
  (testing "Explode Space Separated Nil Content Throws Exception"
    (is (thrown-with-msg? IllegalArgumentException #"content cannot be nil" (explode " " nil)))
    (is (thrown-with-msg? IllegalArgumentException #"content cannot be nil" (explode " " nil 1)))))


(deftest implode-test
  (testing "Can Implode Array Of String"
    #_{:clj-kondo/ignore [:inline-def]}
    (def array ["Hello" "world"])
    (is (= (implode " " array) "Hello world"))
    (is (= (implode "  " array) "Hello  world"))
    (is (= (implode "" array) "Helloworld"))
    (is (= (implode nil array) "Helloworld"))
    (is (= (implode " wkwkwk " array) "Hello wkwkwk world"))
    (is (= (implode "\n" array) "Hello\nworld")))
  (testing "Can Implode Array Of Int"
    #_{:clj-kondo/ignore [:inline-def]}
    (def array [1 2 3 4 5])
    (is (= (implode ":. " array) "1:. 2:. 3:. 4:. 5")))
  (testing "Can Implode Empty Array"
    #_{:clj-kondo/ignore [:inline-def]}
    (def array [])
    (is (= (implode " " array) ""))))

(ns pehape.string.implode-test
  (:require [clojure.test :refer [deftest is testing]]
            [pehape.core :refer [implode]]))


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
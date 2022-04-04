(ns pehape.string.implode
  (:require [clojure.string :as str]))


(defn implode
  "The Implode function joins array elements with a string.

   Arguments:
   - `separator` Specifies what to put between the array elements.
   - `array` The array to join to a string.
   Returns:
   - A string from elements of an array"
  {:added "0.0.1"}
  [separator array]
  (str/join separator array))
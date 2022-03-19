(ns pehape.string
  (:require [clojure.string :as str]))


(defn explode
  "The Explode function breaks a string into an array.
   
   Arguments:
   - `separator` Specifies where to break the string.
   - `content` The string to split.
   - `limit` Specifies the number of array elements to return.
   If limit is greater than 0, then an array with a maximum of limit
   elements is returned. If limit is less than 0 then an array excluding
   the last -limit elements is returned. If limit is 0 then an array with
   one element is returned.
   Returns:
   - An array of strings."
  {:added "0.0.1"}
  ([separator content]
   (cond
     (or (nil? separator)
         (= separator "")) (throw (IllegalArgumentException. "separator cannot be empty"))
     (nil? content) (throw (IllegalArgumentException. "content cannot be empty"))
     :else (str/split content (re-pattern separator))))
  ([separator content limit]
   (cond
     (or (nil? separator)
         (= separator "")) (throw (IllegalArgumentException. "separator cannot be empty"))
     (nil? content) (throw (IllegalArgumentException. "content cannot be empty"))
     (> limit 0) (take limit (str/split content (re-pattern separator)))
     (< limit 0) (drop-last (- limit) (str/split content (re-pattern separator)))
     :else (take 1 (str/split content (re-pattern separator))))))


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
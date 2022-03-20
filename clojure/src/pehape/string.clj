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
         (= separator "")) (throw (IllegalArgumentException. "separator cannot be nil or empty"))
     (nil? content) (throw (IllegalArgumentException. "content cannot be nil"))
     :else (str/split content (re-pattern separator))))
  ([separator content limit]
   (cond
     (or (nil? separator)
         (= separator "")) (throw (IllegalArgumentException. "separator cannot be nil or empty"))
     (nil? content) (throw (IllegalArgumentException. "content cannot be nil"))
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


(defn levenshtein
  "The Levenshtein function returns the Levenshtein distance between two strings.
   
   Arguments:
   - `str1` First string to compare.
   - `str2` Second string to compare.
   - `insert-cost` The cost of inserting a character.
   - `replace-cost` The cost of replacing a character.
   - `delete-cost` The cost of deleting a character.
   Returns:
   - The Levenshtein distance between two argument strings.
   Exceptions:
   - IllegalArgumentException when `str1` or `str2` is null.
   - IllegalArgumentException when `insert-cost`, `replace-cost`, or `delete-cost` is negative"
  {:added "0.0.1"}
  ([str1 str2]
   (cond
     (nil? str1) (throw (IllegalArgumentException. "str1 cannot be nil"))
     (nil? str2) (throw (IllegalArgumentException. "str2 cannot be nil"))
     (zero? (count str1)) (count str2)
     (zero? (count str2)) (count str1)
     :else
     (let [p1 (atom (into [] (range (+ 1 (count str2)))))
           p2 (atom [])
           swapped (atom false)]
       (doseq [i1 (range (count str1))]
         (if @swapped
           (swap! p1 #(assoc % 0 (+ 1 (first @p2))))
           (swap! p2 #(assoc % 0 (+ 1 (first @p1)))))
         (doseq [i2 (range (count str2))]
           (let [c0 (atom (+ (if @swapped
                               (nth @p2 i2)
                               (nth @p1 i2))
                             (if (= (nth str1 i1)
                                    (nth str2 i2)) 0 1)))
                 c1 (atom (+ 1 (if @swapped
                                 (nth @p2 (+ 1 i2))
                                 (nth @p1 (+ 1 i2)))))
                 c2 (atom (+ 1 (if @swapped
                                 (nth @p1 i2)
                                 (nth @p2 i2))))]
             (when (< @c1 @c0) (swap! c0 (fn [_] @c1)))
             (when (< @c2 @c0) (swap! c0 (fn [_] @c2)))
             (if @swapped
               (swap! p1 (fn [n] (assoc n (+ 1 i2) @c0)))
               (swap! p2 (fn [n] (assoc n (+ 1 i2) @c0))))))
         (swap! swapped (fn [n] (not n))))
       (if @swapped
         (nth @p2 (count str2))
         (nth @p1 (count str2))))))
  ([str1 str2 insert-cost replace-cost delete-cost]
   (cond
     (nil? str1) (throw (IllegalArgumentException. "str1 cannot be nil"))
     (nil? str2) (throw (IllegalArgumentException. "str2 cannot be nil"))
     (neg? insert-cost) (throw (IllegalArgumentException. "insert-cost cannot be a negative value"))
     (neg? replace-cost) (throw (IllegalArgumentException. "replace-cost cannot be a negative value"))
     (neg? delete-cost) (throw (IllegalArgumentException. "delete-cost cannot be a negative value"))
     (zero? (count str1)) (* insert-cost (count str2))
     (zero? (count str2)) (* delete-cost (count str1))
     :else
     (let [p1 (atom (into [] (map #(* insert-cost %)
                                  (range (+ 1 (count str2))))))
           p2 (atom [])
           swapped (atom false)]
       (doseq [i1 (range (count str1))]
         (if @swapped
           (swap! p1 #(assoc % 0 (+ delete-cost (first @p2))))
           (swap! p2 #(assoc % 0 (+ delete-cost (first @p1)))))
         (doseq [i2 (range (count str2))]
           (let [c0 (atom (+ (if @swapped
                               (nth @p2 i2)
                               (nth @p1 i2))
                             (if (= (nth str1 i1) (nth str2 i2))
                               0
                               replace-cost)))
                 c1 (atom (+ delete-cost
                             (if @swapped
                               (nth @p2 (+ 1 i2))
                               (nth @p1 (+ 1 i2)))))
                 c2 (atom (+ insert-cost
                             (if @swapped
                               (nth @p1 i2)
                               (nth @p2 i2))))]
             (when (< @c1 @c0) (swap! c0 (fn [_] @c1)))
             (when (< @c2 @c0) (swap! c0 (fn [_] @c2)))
             (if @swapped
               (swap! p1 (fn [n] (assoc n (+ 1 i2) @c0)))
               (swap! p2 (fn [n] (assoc n (+ 1 i2) @c0))))))
         (swap! swapped (fn [n] (not n))))
       (if @swapped
         (nth @p2 (count str2))
         (nth @p1 (count str2)))))))
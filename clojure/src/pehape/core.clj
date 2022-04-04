(ns pehape.core
  (:require [pehape.string.explode :as str-explode]
            [pehape.string.implode :as str-implode]
            [pehape.string.levenshtein :as str-levenshtein]))


(def explode str-explode/explode)
(def implode str-implode/implode)
(def levenshtein str-levenshtein/levenshtein)
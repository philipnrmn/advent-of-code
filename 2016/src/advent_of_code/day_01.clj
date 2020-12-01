(ns advent-of-code.day-01
  (:require [clojure.string :as str]))

(defn scan
  [input]
  (re-seq #"\d+|L|R" input))

(scan "R2 L3")

(defn lex
  [tokens]
  (flatten
    (map 
      (fn [token]
        (let [cmd (read-string token)]
          (if (number? cmd)
            (repeat cmd "1")
            token))) tokens)))

(lex '("R" "2" "L" "3"))
(lex (scan "R2 L3"))

(defn pivot
  [heading bearing]
  (mod (+ heading bearing) 360))

(pivot 270 90)
(pivot 0 -90)

(defn move
  [[x y] distance heading]
  (cond
    (= heading 0)
    [x (- y distance)]
    (= heading 90)
    [(+ x distance) y]
    (= heading 180)
    [x (+ y distance)]
    (= heading 270)
    [(- x distance) y]))

(move [0 0] 5 0)
(move [0 0] 5 90)
(move [0 0] 5 180)
(move [0 0] 5 270)

(defn move-to
  [input]
  (loop [loc [0 0] heading 0 tokens (scan input)]
    (if (= 0 (count tokens))
      loc
      (let [cmd (first tokens)]
        (cond
          (= cmd "L")
          (recur loc (pivot heading -90) (rest tokens))
          (= cmd "R")
          (recur loc (pivot heading 90) (rest tokens))
          :else 
          (recur (move loc (Integer/parseInt cmd) heading) heading (rest tokens)))))))

(move-to "R2. L3")
(move-to "R2, R2, R2")
(move-to "R5, L5, R5, R3")

(move-to "L1, L5, R1, R3, L4, L5, R5, R1, L2, L2, L3, R4, L2, R3, R1, L2, R5, R3, L4, R4, L3, R3, R3, L2, R1, L3, R2, L1, R4, L2, R4, L4, R5, L3, R1, R1, L1, L3, L2, R1, R3, R2, L1, R4, L4, R2, L189, L4, R5, R3, L1, R47, R4, R1, R3, L3, L3, L2, R70, L1, R4, R185, R5, L4, L5, R4, L1, L4, R5, L3, R2, R3, L5, L3, R5, L1, R5, L4, R1, R2, L2, L5, L2, R4, L3, R5, R1, L5, L4, L3, R4, L3, L4, L1, L5, L5, R5, L5, L2, L1, L2, L4, L1, L2, R3, R1, R1, L2, L5, R2, L3, L5, L4, L2, L1, L2, R3, L1, L4, R3, R3, L2, R5, L1, L3, L3, L3, L5, R5, R1, R2, L3, L2, R4, R1, R1, R3, R4, R3, L3, R3, L5, R2, L2, R4, R5, L4, L3, L1, L5, L1, R1, R2, L1, R3, R4, R5, R2, R3, L2, L1, L5")
; [107 -146]

(defn move-to-and-trace
  [input]
  (loop [loc [0 0] heading 0 tokens (lex (scan input)) history [[0 0]]]
    (if (= 0 (count tokens))
      history
      (let [cmd (first tokens)]
        (cond
          (= cmd "L")
          (recur loc (pivot heading -90) (rest tokens) history)
          (= cmd "R")
          (recur loc (pivot heading 90) (rest tokens) history)
          :else 
          (let [new-loc (move loc (Integer/parseInt cmd) heading)]
          (recur new-loc heading (rest tokens) (conj history new-loc))))))))

(move-to-and-trace "R2, L3")

(defn in-seq?
  [haystack needle]
  (some #(= needle %) haystack))

(defn find-first-recurrence
  [history]
  (loop [locations history]
    (let [here (nth locations 0) there (next locations)]
      (if (in-seq? there here)
        here
        (recur there)))))

(find-first-recurrence [1 2 3 2])
(find-first-recurrence (move-to-and-trace "R8, R4, R4, R8"))

(find-first-recurrence (move-to-and-trace "L1, L5, R1, R3, L4, L5, R5, R1, L2, L2, L3, R4, L2, R3, R1, L2, R5, R3, L4, R4, L3, R3, R3, L2, R1, L3, R2, L1, R4, L2, R4, L4, R5, L3, R1, R1, L1, L3, L2, R1, R3, R2, L1, R4, L4, R2, L189, L4, R5, R3, L1, R47, R4, R1, R3, L3, L3, L2, R70, L1, R4, R185, R5, L4, L5, R4, L1, L4, R5, L3, R2, R3, L5, L3, R5, L1, R5, L4, R1, R2, L2, L5, L2, R4, L3, R5, R1, L5, L4, L3, R4, L3, L4, L1, L5, L5, R5, L5, L2, L1, L2, L4, L1, L2, R3, R1, R1, L2, L5, R2, L3, L5, L4, L2, L1, L2, R3, L1, L4, R3, R3, L2, R5, L1, L3, L3, L3, L5, R5, R1, R2, L3, L2, R4, R1, R1, R3, R4, R3, L3, R3, L5, R2, L2, R4, R5, L4, L3, L1, L5, L1, R1, R2, L1, R3, R4, R5, R2, R3, L2, L1, L5"))
; [123 -3]

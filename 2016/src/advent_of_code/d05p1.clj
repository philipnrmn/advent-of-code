(ns advent-of-code.d04p1)

(import 'java.security.MessageDigest
        'java.math.BigInteger)

(defn md5 [s]
  (let [algorithm (MessageDigest/getInstance "MD5")
        size (* 2 (.getDigestLength algorithm))
        raw (.digest algorithm (.getBytes s))
        sig (.toString (BigInteger. 1 raw) 16)
        padding (apply str (repeat (- size (count sig)) "0"))]
    (str padding sig)))

(md5 "abc3231929")

; for full digest support, use clj-digest: https://github.com/tebeka/clj-digest


(defn matching?
  [s]
  (= (.indexOf s "00000") 0))

(matching? "000001")
(matching? "100000")

(def input "wtnhxymk")

(def hashes
  (map
    #(md5 (str input %))
    (iterate inc 1)))

(def matching-hashes
  (take 8
    (filter matching? hashes)))
; 0000027b9705c7e6fa3d4816c490bbfd
; 00000468c8625d85571d250737c47b5a
; 0000013e3293b49e4c78a5b43b21023b
; 0000040bbe4509b48041007dec6123bd
; 00000b11810477f9e49840991fb2151e
; 00000cc461c8945671046cf632be4473
; 000007c1da6865df78b2c0addf28913d
; 00000700ce8beb0a8ffc83fa9986d577

(def password
  (apply str (map #(.charAt % 5) matching-hashes)))
; 2414bc77

(defn substitute
  [arr x n]
  (concat
    (subvec arr 0 n)
    [x]
    (subvec arr (inc n) (count arr))))

(substitute [1 2 3 4 5] :a 2)
(substitute [1 2 3 4 5] :a 4)

(some nil? [1 2 nil 3])

; _37d60f

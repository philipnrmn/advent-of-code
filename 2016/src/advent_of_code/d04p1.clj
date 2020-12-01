(ns advent-of-code.d04p1)

(def inputs
  (clojure.string/split
    (slurp "./input/4.txt")
    #"\n"))

(take 5 inputs)

(defn strip
  [s]
  (apply str
    (filter #(not (= \- %)) s)))

(strip "aa-bb-cc")

(defn count-chars
  [c s]
  (count
    (filter #(= c %) s)))

(count-chars \a "aabbccaa")
(count-chars \o "notarealroom")

(defn checksum
  [encrypted-name]
  (let [letters (set (.toCharArray encrypted-name))]
    (apply str
     (take 5
       (reverse
         (sort-by #(count-chars % encrypted-name) (reverse (sort letters))))))))

(checksum "aaaaabbbzyx")
(checksum "zyxaaaaabbb")
(checksum "abcdefgh")
(checksum "notarealroom")

(defn split
  [input]
  (let [[_ nam id sum] 
        (re-find #"([a-z]+)(\d+)\[(\w+)\]" (strip input))]
    {:name nam :id id :sum sum}))

(split "aaaaa-bbb-z-y-x-123[abxyz]")
(split "a-b-c-d-e-f-g-h-987[abcde]")
(split "not-a-real-room-404[oarel]")

(defn valid?
  [input]
  (let [{nam :name sum :sum} (split input)]
    (= sum (checksum nam))))

(valid? "aaaaa-bbb-z-y-x-123[abxyz]")
(valid? "a-b-c-d-e-f-g-h-987[abcde]")
(valid? "not-a-real-room-404[oarel]")
(valid? "totally-real-room-200[decoy]")

(def valid-inputs (filter valid? inputs))
(def valid-input-ids (map #(Integer/parseInt (re-find #"\d+" %)) valid-inputs))

(apply + valid-input-ids)
; 409147

(defn shift-char
  [c n]
  (let [a (int \a) i (- (int c) a)]
    (char (+ (mod (+ i n) 26) a))))

(shift-char \a 25)

(defn shift-str
  [s n]
  (apply str (map #(shift-char % n) s)))

(shift-str "qzmtzixmtkozyivhz" 343)

(def decoded-names
  (map
    (fn [{nam :name id :id}]
      {:name (shift-str nam (Integer/parseInt id)) :id id})
    (map split valid-inputs)))

(filter
  (fn [{nam :name id :id}]
    (> (.indexOf nam "north") -1))
  decoded-names)
; 991
      



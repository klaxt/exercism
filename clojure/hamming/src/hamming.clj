(ns hamming)

(defn validate [strand1 strand2]
  (= (count strand1) (count strand2))
  )

(defn distance [strand1 strand2]
  (if (validate strand1 strand2)
    (count (filter false? (map = strand1 strand2)))
    nil
    )
  )

(defn distanceV1 [strand1 strand2]
  (if (not (validate strand1 strand2))
    nil
    (loop [i 0 r []]
      (if (< i (count strand2))
        (let [e1 (nth strand1 i) e2 (nth strand2 i)]
          (if (= e1 e2)
            (recur (inc i) r)
            (recur (inc i) (conj r e1))
            )
          )
        (count r)
        )
      )
    )
  )

(defn fibonacci [count]
  (loop [res [0 1]]
    (if (>= (count res) count)
      res
      (recur
        (conj res
              (+'
                (last res)
                (last (butlast res))
                )
              )
        )
      )
    )
  )
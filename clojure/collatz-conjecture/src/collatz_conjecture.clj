(ns collatz-conjecture)

(defn calculate [num]
  (if (even? num)
    (/ num 2)
    (inc (* num 3))
  )
)

(defn collatz [num]
  (cond
    (= num 0) (throw (Exception. "zero is an error"))
    (< num 0) (throw (Exception. "negative value is an error"))
    :else (->> num
               (iterate calculate)
               ;https://clojure.org/guides/learn/functions#_anonymous_function_syntax
               (take-while #(not= 1 %))
               (count)
               )
    )
  )

(defn collatzV1 [num]
  (cond
    (= num 0) (throw (Exception. "zero is an error"))
    (< num 0) (throw (Exception. "negative value is an error"))
    :else
    (loop [value num i 0]
      (println value " " i)
      (if (not (= value 1))
        (if (even? value)
          (recur (/ value 2) (inc i))
          (recur (+ (* value 3) 1) (inc i))
          )
          i
        )
      )
    )
  )

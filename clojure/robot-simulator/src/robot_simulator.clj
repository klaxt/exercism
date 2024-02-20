(ns robot-simulator)

;https://clojure.org/reference/datatypes
(defrecord Robot [coordinates bearings])

(defn robot [coordinates bearing]
  ;(Robot. coordinates bearing)
  {:coordinates coordinates :bearing bearing}
)

(defn advance[robot]
  (cond
    (= (:bearing robot) :north) (update-in robot [:coordinates :y] inc)
    (= (:bearing robot) :east) (update-in robot [:coordinates :x] inc)
    (= (:bearing robot) :south) (update-in robot [:coordinates :y] dec)
    (= (:bearing robot) :west) (update-in robot [:coordinates :x] dec)
    )
  ;(update robot :coordinates)
  )

(defn turn-right [bearing]
  (cond
    (= bearing :west) :north
    (= bearing :north) :east
    (= bearing :east) :south
    (= bearing :south) :west
    )
)

(defn turn-left [bearing]
  (->> bearing
      ;(iterate turn-right)
      ;(take-nth 3)
    turn-right
    turn-right
    turn-right
    )
)

(defn move[step robot]
  (cond
    (= \A step) (advance robot)
    (= \L step) (update robot :bearing turn-left)
    (= \R step) (update robot :bearing turn-right)
    :else (throw (Exception. "invalid step"))
    )
  )

(defn simulate [path robot]
  (loop [i 0 r robot]
    (if (< i (count path))
      (recur (inc i) (move (nth path i) r))
      r
      )
    )
  )


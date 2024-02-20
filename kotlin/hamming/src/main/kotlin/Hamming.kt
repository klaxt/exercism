object Hamming {

    fun compute(leftStrand: String, rightStrand: String): Int {
        if (leftStrand.length != rightStrand.length) {
            throw IllegalArgumentException("left and right strands must be of equal length")
        }
        return leftStrand.zip(rightStrand).fold(0) {
            acc, it -> acc + it.diff()
        }
    }

    private fun Pair<Char, Char>.diff() : Int = if (first == second) 0 else 1
}


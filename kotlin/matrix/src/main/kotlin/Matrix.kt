class Matrix(private val matrixAsString: String) {
    private val array = matrixAsString.split('\n').map {
        it.split(' ').map { it.toInt() }
    }

    fun column(colNr: Int): List<Int> = array.map { it[colNr - 1] }

    fun row(rowNr: Int): List<Int> = array[rowNr - 1]
}

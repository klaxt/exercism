import java.util.stream.IntStream

class Squares(val n: Int) {
    //TODO: implement proper constructor

    fun sumOfSquares(): Int =
        (1..n).reduce { acc, x -> acc + x.square() }

    fun squareOfSum(): Int = (1..n).sum().square()

    fun difference(): Int = squareOfSum() - sumOfSquares()

    fun Int.square() = this * this
}

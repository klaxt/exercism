import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond(
    date: LocalDateTime
) {
    constructor(d: LocalDate) : this(d.atStartOfDay())

    val date: LocalDateTime = date.plusSeconds(GIGASECOND)

    companion object {
        const val GIGASECOND = 1000000000L
    }
}

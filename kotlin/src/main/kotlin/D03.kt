fun main() {
    val lines = readLinesFromResource("d03").sorted()

    var gamma = 0
    var epsylon = 0
    for (i in 0 until lines[0].length) {
        val ones = lines.count { it[i] == '1' }
        val zeros = lines.count { it[i] == '0' }
        gamma = gamma shl 1
        epsylon = epsylon shl 1
        if (ones > zeros) gamma++ else epsylon++
    }

    var oxyLines = lines.toList()
    var i = 0
    while (oxyLines.size > 1) {
        val ones = oxyLines.count { it[i] == '1' }
        val zeros = oxyLines.count { it[i] == '0' }
        oxyLines = oxyLines.filter {
            if (ones >= zeros) {
                it[i] == '1'
            } else {
                it[i] == '0'
            }
        }
        i++
    }

    var co2Lines = lines.toList()
    i = 0
    while (co2Lines.size > 1) {
        val ones = co2Lines.count { it[i] == '1' }
        val zeros = co2Lines.count { it[i] == '0' }
        co2Lines = co2Lines.filter {
            if (ones < zeros) {
                it[i] == '1'
            } else {
                it[i] == '0'
            }
        }
        i++
    }

    println("part1 " + gamma * epsylon)
    println("part2 " + oxyLines[0].toInt(2) * co2Lines[0].toInt(2))
}



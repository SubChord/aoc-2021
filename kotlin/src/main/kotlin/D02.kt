fun main() {
    val lines = readLinesFromResource("d02")

    var x1 = 0
    var y1 = 0
    lines.forEach {
        val split = it.split(" ")
        val v = split[1].toInt()

        when (split[0]) {
            "forward" -> x1 += v
            "up" -> y1 -= v
            "down" -> y1 += v
        }
    }

    println("Part 1: " + (x1 * y1))

    var x2 = 0
    var y2 = 0
    var aim = 0
    lines.forEach {
        val split = it.split(" ")
        val v = split[1].toInt()

        when (split[0]) {
            "forward" -> {
                x2 += v
                y2 += v * aim
            }
            "up" -> aim -= v
            "down" -> aim += v
        }
    }

    println("Part 2: " + (x2 * y2))
}



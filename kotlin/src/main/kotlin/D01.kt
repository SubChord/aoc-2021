fun main() {
    val lines = readLinesFromResource("d01")

    val ints = lines.map { it.toInt() }
    ints.mapIndexed { index, i ->
        val get = ints.getOrNull(index + 1)
        if ((get ?: 0) > i) {
            1
        } else {
            0
        }
    }.sum().let { println("part 1: $it") }

    ints.takeLast(ints.size - 3)
        .mapIndexed { index, _ ->
            val i = index + 3
            val s1 = ints.slice(i - 3 until i).sum()
            val s2 = ints.slice(i - 2..i).sum()
            if (s2 > s1) 1 else 0
        }
        .sum().let { println("part 2: $it") }
}



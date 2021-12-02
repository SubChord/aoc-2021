fun readLinesFromResource(resourceName: String): List<String> {
    val inputStream = object {}.javaClass.getResourceAsStream(resourceName)
    return inputStream.bufferedReader().useLines { it.toList() }
}

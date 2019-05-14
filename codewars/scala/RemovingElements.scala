object Kata {
  def removeEveryOther[T](list: List[T]): List[T] = {
    list.zipWithIndex.filter{ case (datum, index) => index % 2 == 0 }.map(_._1)
  }
}

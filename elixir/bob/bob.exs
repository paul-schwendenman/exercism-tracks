defmodule Bob do
  def hey(input) do
    yelling = yelling?(input)
    question = question?(input)

    cond do
      question && yelling -> "Calm down, I know what I'm doing!"
      question -> "Sure."
      yelling -> "Whoa, chill out!"
      silence?(input) -> "Fine. Be that way!"
      true -> "Whatever."
    end
  end

  defp yelling?(normal) do
    upper = String.upcase(normal)
    lower = String.downcase(normal)

    (upper == normal) && (upper != lower)
  end

  defp silence?(input) do
    String.trim(input) == ""
  end

  defp question?(input) do
    String.ends_with?(input, "?")
  end
end

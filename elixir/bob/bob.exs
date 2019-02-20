defmodule Bob do
  def hey(input) do
    is_yelling = is_yelling?(input)
    is_question = is_question?(input)

    cond do
      is_question && is_yelling -> "Calm down, I know what I'm doing!"
      is_question -> "Sure."
      is_yelling -> "Whoa, chill out!"
      is_silence?(input) -> "Fine. Be that way!"
      true -> "Whatever."
    end
  end

  defp is_yelling?(normal) do
    upper = String.upcase(normal)
    lower = String.downcase(normal)

    (upper == normal) && (upper != lower)
  end

  defp is_silence?(input) do
    String.trim(input) == ""
  end

  defp is_question?(input) do
    String.ends_with?(input, "?")
  end
end

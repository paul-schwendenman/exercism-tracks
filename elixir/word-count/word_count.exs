defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence
    |> String.downcase
    |> String.split(" ")
    |> Enum.map(fn x -> {x, 1} end)
    |> Map.new
  end
end

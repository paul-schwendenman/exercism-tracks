defmodule Bob do
  def hey(input) do
    cond do
      String.match?(input, ~r/^\s*$/) -> "Fine. Be that way!"
      String.match?(input, ~r/^[\d ,]+$/) -> "Whatever."
      String.match?(input, ~r/^[A-Z' ]+\?$/) -> "Calm down, I know what I'm doing!"
      String.match?(input, ~r/^.+\?$/) -> "Sure."
      String.match?(input, ~r/^[^a-z]+!?$/) -> "Whoa, chill out!"
      true -> "Whatever."
    end
  end
end

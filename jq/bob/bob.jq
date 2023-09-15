.heyBob
| sub("\\s*$"; "")
| test("\\?$") as $is_question
| test("^[^a-z]*[A-Z][^a-z]*$") as $is_yelling
| test("^\\s*$") as $is_empty
|   if $is_empty then
        "Fine. Be that way!"
    elif $is_yelling and $is_question then
        "Calm down, I know what I'm doing!"
    elif $is_yelling then
        "Whoa, chill out!"
    elif $is_question then
        "Sure."
    else
        "Whatever."
    end

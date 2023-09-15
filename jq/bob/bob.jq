def rstrip: sub("\\s+$"; "");
def is_question: test("\\?$");
def is_empty: test("^$");
def is_yelling: test("^[^a-z]*[A-Z][^a-z]*$");

.heyBob |
rstrip |
{
    empty: (is_empty),
    yelling: (is_yelling),
    question: (is_question)
} |
if .empty then
    "Fine. Be that way!"
elif .yelling and .question then
    "Calm down, I know what I'm doing!"
elif .yelling then
    "Whoa, chill out!"
elif .question then
    "Sure."
else
    "Whatever."
end

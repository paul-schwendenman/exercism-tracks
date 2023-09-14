# Given a certain command, help the Chatbot recognize whether
# the command is valid or not.
# - valid if "Chatbot" is at the start of the input
#
# input: {string} a command
# output: {boolean} whether or not the command is valid
#
def is_valid_command:
  # (.[0:7] | ascii_downcase ) == "chatbot"
  . | test("^chatbot[^a-z0-9_]"; "i")
;

# Given a certain message, help the Chatbot get rid of all the
# "encrypted" emojis throught the message.
# - an "encrypted emoji" is the string "emoji" followed by digits
#
# input: {string} message
# output: {string} the message without the emojis
def remove_emoji:
  . | sub("emoji[0-9]+"; ""; "g")
;

# Given a certain phone number, help the Chatbot recognize
# whether it is in the correct format.
# - a phone number has the form "(+NN) NNN-NNN-NNN"
#   where N is a digit
#
# input: {string} number
# output: {string} the Chatbot response to the phone validation
def check_phone_number:
  . | scan("[-+()0-9]+ ?[-0-9]+") | if test("\\(\\+[0-9]{2}\\) [0-9]{3}(-[0-9]{3}){2}") then "Thanks! Your phone number is OK."
  else "Oops, it seems like I can't reach out to \(.)."
  end
;

# Given a certain response from the user, help the Chatbot get
# only the website domains
# - a domain consists of sequences of word characters joined by dots
#
# input: {string} userInput
# output: {array} all the domains in the input
def get_domains:
  . | [ scan("[A-Za-z]+\\.[A-Za-z.]+") ]
;

# Greet the user using their name
# - the phrase "my name is Something" must be in the input
#
# input: {string} sentence with name clause
# output: {string} greeting from the Chatbot
def nice_to_meet_you:
  . | capture("my name is (?<name>[-A-Za-z]+)"; "i") | "Nice to meet you, \(.name)"
;

# Perform very simple CSV parsing
# - fields are separated by a comma and optional whitespace
#
# input: {string} comma-separated row
# output: {array} fields
def parse_csv:
  . | split(",") | map(scan("[A-Za-z0-9]+"))
;

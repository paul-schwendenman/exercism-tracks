def raindrop_sound:
    . | [
        if . % 3 == 0 then "Pling" else "" end,
        if . % 5 == 0 then "Plang" else "" end,
        if . % 7 == 0 then "Plong" else "" end
    ] | add
;

.number | if raindrop_sound != "" then raindrop_sound else . end

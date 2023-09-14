.number | if (. % 3 == 0) or (. % 5 == 0) or (. % 7 == 0) then ([
    if . % 3 == 0 then "Pling" else "" end,
    if . % 5 == 0 then "Plang" else "" end,
    if . % 7 == 0 then "Plong" else "" end
] | add) else . end

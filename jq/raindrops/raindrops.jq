.number | if . % 3 == 0 then
    if . % 5 == 0 then
        if . % 7 == 0 then "PlingPlangPlong"
        else "PlingPlang"
        end
    elif . % 7 == 0 then "PlingPlong"
    else "Pling"
    end
elif . % 5 == 0 then
    if . % 7 == 0 then "PlangPlong"
    else "Plang"
    end
elif . % 7 == 0 then "Plong"
else .
end

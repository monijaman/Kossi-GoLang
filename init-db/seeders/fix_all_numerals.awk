BEGIN {
    # Map ASCII to Bengali digits
    ascii_to_bengali["0"] = "०"
    ascii_to_bengali["1"] = "१"
    ascii_to_bengali["2"] = "२"
    ascii_to_bengali["3"] = "३"
    ascii_to_bengali["4"] = "४"
    ascii_to_bengali["5"] = "५"
    ascii_to_bengali["6"] = "६"
    ascii_to_bengali["7"] = "७"
    ascii_to_bengali["8"] = "८"
    ascii_to_bengali["9"] = "९"
}

{
    # Process lines that are map entries: "KEY": "VALUE",
    if ($0 ~ /^\s*"[^"]*":\s*".*",?\s*$/) {
        # Split on ": " to get key and value parts
        # Find the position of ": "
        idx = index($0, "\": \"")
        if (idx > 0) {
            # Extract key part (from start to ": ")
            key_part = substr($0, 1, idx + 2)  # includes ": "
            # Extract rest (value and closing quotes/comma)
            rest = substr($0, idx + 4)  # skip the opening quote after ": "
            
            # Find the closing quote of the value
            close_idx = match(rest, /"/)
            if (close_idx > 0) {
                value = substr(rest, 1, close_idx - 1)
                closing_part = substr(rest, close_idx)
                
                # Convert digits in value
                for (i = 0; i <= 9; i++) {
                    gsub(i, ascii_to_bengali[i ""], value)
                }
                
                # Reconstruct the line
                $0 = key_part "\"" value closing_part
            }
        }
    }
    print $0
}

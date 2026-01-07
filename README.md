# REGEX IN GO
## Regular expression can be defined as pattern or template. Can be used for search and can replace input validation

## flags (example (?i))
- `i` (ignore) makes regex case insensetive
- `g` (global) searches for matches in all text(by default only first match)
- `m` (multiline) affects anchors behaviour `^` and `$` (start and end of a string) making regex to apply to multiple lines
- `s` (single-line) treats text as one-line string

## metacharacters

### character classes matches specific set of characters

- `\d` - numeric 
- `\w` - word (alphanumeric + underscore)
- `\s` - matches any whitespace character


> **_NOTE:_**  capitalizing char classes will make them work the opposite way. \D will search for non-digit chars

### Negated character set

- `^` (all but) example: all but vowes: `[^aeiouAEIOU\s]`

### ranges
- `[0-9]` -- matches character from 0 to 9
- `[a-z]`
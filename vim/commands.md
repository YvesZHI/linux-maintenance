# Commands

target word: aaa

given word: 111

### Replace the first target word with the given word at current line
```
:s/aaa/111
```

### Replace all of target words with the given word at current line 
```
:s/aaa/111/g
```

### Replace the first target word of each line with the given word
```
:%s/aaa/111
```

### Replace all of the target words of each line with the given word
```
:%s/aaa/111/g
```

### Delete all lines containing the target word
```
:g/aaa/d
```

### Replace all of the target words of lines containing 'bbb' with the given word
```
:g/bbb/%s/aaa/111/g
```


# Regular expression

target word: aaa

greedy: `.*`

non-greedy: `.\{-}`

### get the last target word
```
.*\zsaaa
```

### get the content from the beginning to the second target word
```
\(.\{-}aaa\)\{2}
```

### get the content from the beginning to the second target word(excluding)
```
\(.\{-}\zeaaa\)\{2}
```

### get the content from the second target word to the end
```
\(aaa.\{-}\)\{2}
```

### get the second target word
```
\(.\{-}\zsaaa\)\{2}
```

### get the content from the first target word to the second target word
```
\(aaa.\{-}\)\{2}
```

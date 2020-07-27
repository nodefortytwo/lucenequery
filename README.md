# lucenequery
 Simple library for composing lucene queries
 
 ## Example
```go
q = lq.New(lq.Or{
  []lq.Clause{
    lq.And{
      []lq.Clause{
        lq.Term{"title", "foo bar"},
        lq.Term{"body", "quick fox"},
      },
    },
    lq.Term{"title", "fox"},
  },
}).String()
```

would create

```
((title:"foo bar" AND body:"quick fox") OR title:fox)
```

more complex examples in the tests


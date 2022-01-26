module github.com/robertDurst/GoExplore

go 1.17

require (
    "github.com/robertDurst/GoExplore/interpreter/lexar" v0.0.0
    "github.com/robertDurst/GoExplore/interpreter/tokenizer" v0.0.0
)

replace (
    "github.com/robertDurst/GoExplore/interpreter/lexar" v0.0.0  => "./interpreter/lexar"
    "github.com/robertDurst/GoExplore/interpreter/tokenizer" v0.0.0  => "./interpreter/tokenizer"
)
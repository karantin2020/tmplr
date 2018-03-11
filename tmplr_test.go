package main 

//go:generate ./tmplr -v
//go:generate ./tmplr -h
//go:generate ./tmplr -t ./tmpls/SGBson -g
//go:generate ./tmplr -t SGBson -o ./tmp/IntStr str:Int lst:nums
//go:generate ./tmplr -t SGBson --append -o ./tmp/IntStr str:CondKind lst:conds
//go:generate ./tmplr -t SGBson --append -o ./tmp/IntStr str:CondDo lst:cords

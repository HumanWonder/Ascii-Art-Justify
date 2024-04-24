echo go run . --align right something standard
go run . --align right something standard

echo go run . --align=right left standard
go run . --align=right left standard

echo go run . --align=left right standard
go run . --align=left right standard

echo go run . --align=center hello shadow
go run . --align=center hello shadow

echo go run . --align=justify "1 two 4" shadow
go run . --align=justify "1 two 4" shadow

echo go run . --align=right 23/32 standard
go run . --align=right 23/32 standard

echo go run . --align=right ABCabc123 thinkertoy
go run . --align=right ABCabc123 thinkertoy

echo go run . --align=center "#$%&\"" thinkertoy
go run . --align=center "#$%&\"" thinkertoy

echo go run . --align=left '23Hello World!' standard
go run . --align=left '23Hello World!' standard

echo go run . --align=justify "HELLO there HOW are YOU?!" thinkertoy
go run . --align=justify "HELLO there HOW are YOU?!" thinkertoy

echo go run . --align=right "a -> A b -> B c -> C" shadow
go run . --align=right "a -> A b -> B c -> C" shadow

echo go run . --align=right abcd shadow
go run . --align=right abcd shadow

echo go run . --align=center ola standard
go run . --align=center ola standard

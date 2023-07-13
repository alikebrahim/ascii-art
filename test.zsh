#!/usr/bin/env zsh

test() {
    echo "Test (1): go run . --align right "hello" standard"
    go run . --align right "hello" standard 
    echo 

    echo "Test (2): go run . --align=right left standard"
    go run . --align=right left standard
    echo

    echo "Test (3): go run . --align=left right standard"
    go run . --align=left right standard
    echo

    echo "Test (4): go run . --align=center hello shadow "
    go run . --align=center hello shadow 
    echo

    echo "Test (5): go run . --align=justify "1 Two 4" shadow " 
    go run . --align=justify "1 Two 4" shadow 
    echo
    
    echo "Test (6): go run . --align=right 23/32 standard"
    go run . --align=right 23/32 standard
    echo

    echo "Test (7): go run . --align=right ABCabc123 thinkertoy"
    go run . --align=right ABCabc123 thinkertoy
    echo

    echo "Test (8): go run . --align=center "#$%&\"" thinkertoy"
    go run . --align=center "#$%&\"" thinkertoy
    echo

    echo "Test (9): go run . --align=left '23Hello World!' standard"
    go run . --align=left '23Hello World!' standard
    echo

    echo "Test (10): go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy" 
    go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy
    echo

    echo "Test (11): go run . --align=right "a -> A b -> B c -> C" shadow" 
    go rechoun . --align=right "a -> A b -> B c -> C" shadow
    echo

    echo "Test (12): go run . --align=right abcd shadow" 
    echo "***WARNING*** Test requires reducing the window size" 
    go run . --align=right abcd shadow
    echo
    
    echo "Test (13): go run . --align=center ola standard" 
    echo "***WARNING*** Test requires reducing the window size" 
    go run . --align=center ola standard
    echo


    go run . --align=center "***********" shadow
    go run . --align=center "Now Run" standard
    go run . --align=center "Prescribed Tests" standard

}
test
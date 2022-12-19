## TODO

[x] fix floodfill / lookahead to count empty squares each direction
    - updated to the status of the ts-snake floodfill.  still missing the last "wrap-around" dimension

adjust move priority to be okay with moving nearer to smaller snakes (currently avoids larger)

check out if there's any way to print color in go terminal

[x] hook up database connection to store move state for writing tests
    - database is connected, have used it to write some tests to start

write tests for existing funcs

build out logic for wrapped mode and hazard mode


10 PRINT "HELLO, WORLD!"
20 LET A = 10
30 LET B = 20
40 LET C = A + B
50 PRINT "A + B = ", C
60 PRINT "ENTER A NUMBER: "
70 INPUT D
80 IF D > 10 THEN GOTO 110
90 PRINT "THE NUMBER IS LESS THAN OR EQUAL TO 10"
100 GOTO 120
110 PRINT "THE NUMBER IS GREATER THAN 10"
120 LET E = 1 + 2 + 3 + 4 + 5
130 PRINT "SUM OF 1 TO 5 = ", E
140 LET G = 150
150 GOTO 150 + 50
200 PRINT "PROGRAM END"
210 END
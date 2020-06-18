# Golang_Interfaces_task

Task 1 
<br>

Write a program that creates two custom struct types called 'Triangle' and 'Square'
<br>
The 'square' type should be a struct with a field called 'sidelength'
<br>
The 'triangle' type should be a struct with a field called 'height' of type float64 and a field of type 'base' of type float64
<br>
Both types should have function called 'getArea' that returns the calculated area of the square or triangle 
<br>
Add a 'shape' interface that deines a function called 'printArea' This function should calculate the area of the given shape and print it out to the terminal Design the interface so that the 'printArea' function can be called with either a triangle or a square.
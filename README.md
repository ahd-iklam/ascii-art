ASCII Art Program
Overview
Ascii-art is a program that converts a given string into a graphic representation using ASCII characters. The program takes specific parameters from the command line to generate the output in different styles and save it to a specified file.

AUTHOR: Ahmed Malki
Usage
To run the program, use the following format:
go run . [OPTION] [STRING] [BANNER]
Where:

[OPTION] is the flag indicating the output file.
[STRING] is the text you want to convert into ASCII art.
[BANNER] specifies the style of the ASCII art.
Example
To output the string "something" in the "standard" style to a file named fileName.txt, use:
go run . --output=<fileName.txt> something standard
Options
--output=<fileName.txt>: This flag specifies the name of the file where the ASCII art will be saved. The format must be exactly as shown, including the --output= prefix.
Usage Message
If the flag format does not match the specified format, the program will return the following usage message:
Usage: go run . [OPTION] [STRING] [BANNER]
Examples
Here are some examples of how to use the program:

Standard Style:
go run . --output=output.txt HelloWorld standard
This command will generate an ASCII art representation of "HelloWorld" in the standard style and save it to output.txt
Shadow Style:
go run . --output=shadow.txt Goodbye shadow
Thinkertoy Style:
go run . --output=thinkertoy.txt Welcome thinkertoy
This command will generate an ASCII art representation of "Welcome" in the thinkertoy style and save it to thinkertoy.txt.

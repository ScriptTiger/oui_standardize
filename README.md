[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://docs.google.com/forms/d/e/1FAIpQLSfBEe5B_zo69OBk19l3hzvBmz3cOV6ol1ufjh0ER1q3-xd2Rg/viewform)

# oui_standardize
oui_standardize standardizes IEEE MAC address block files (CSV files) containing OUI records by importing an arbitrary number of said files, sorting them by MAC address prefix, and then outputting the results to standard output (stdout). Curly brackets are used as delimiters and multi-line fields are reduced to single lines in order to allow scripts to more easily parse the records with as little overhead as possible.

Usage: `oui_standardize [<file>] [<file>] [<file>]...`

# More About ScriptTiger

For more ScriptTiger scripts and goodies, check out ScriptTiger's GitHub Pages website:  
https://scripttiger.github.io/

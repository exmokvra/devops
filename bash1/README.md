Notes about BASH:

If it's something beyond 15 rows, use a better programming language

To pre load variables, functions and configurations:
~/.bashrc
/etc/profile

Create aliases:
- ALIAS <alias que quero usar>=’<comando real do Linux>’
- Example: ALIAS cl=’clear’

Basic hash commands:
- <variable>=”guilherme”
- ECHO $<variable>
- Or ECHO “$<variable> test << gonna print “<variable> test"
- SED = substitute something inside a file
- ECHO “welcome to the class” | sed “s/class/team”
-- Gonna print “welcome to the team”
- WGET = download something from internet
- TAR = unzip

BASH operators:
String operator - Effect
= - equal
!= - Not equal
< - Less than
> - Greater than
-n s1 - String s1 is not empty
-z s1 - String s1 is empty

Arithmetic comparisons - Effect
-lt - < (less than)
-gt - > (greater than)
-le - <= (less or equals)
-ge - >= (greater or quals)
-eq - == equals
-ne - != not equals

If structure:
“ 
IF [ $s1 = $s2 ]; THEN
	<operation>
ELSE
	<operation2>
FI
“

VI commands:
- i = edit file
- ESC = leave edit mode
- :Q = exit
- :WQ = save and exit

Bash sample code:
- Cd /tmp << goes to folder /tmp
- Mkdir mydir << creates mydir folder
- Cd mydir << goes to mydir folder
- Touch x.txt << creates x.txt file
- Echo “123456” > x.txt << writes 123456 inside file x.txt
- Cat x.txt << shows the file content
- Ps aux | grep ps > x.txt << writes all the result of PS command inside the file
-- With just a > overwrites everything
-- With two >> will make an append to the end of file
- Head x.txt << gets the beginning of file
- Tail x.txt << gets the end of file
- Ps aux | grep ps > x.txt << gets info and throws inside the file

Create the file that will containg code above:
- Chmod +x s.sh << gives running permission to the file
- ./s.sh << runs the file

grammar npconf;

npconf
    : (section | EOL)*
    ;

section
   : sectionheader line*
   | sectionheader section* sectionfooter
   | sectionheader line* sectionfooter
   ;

sectionheader
   : '[' str ']' EOL
   ;

sectionfooter
   : '[' '/' str ']' EOL
   ;

str
   : CHARS
   | STRING
   ;

line
   : stringlist EOL
   | stringlist
   ;

stringlist
   : str (',' str?)*
   ;


CHARS
   : ('A' .. 'Z' | '0' .. '9' | 'a' .. 'z' | '.' | '%' | '"' | '\\' | '/' | '*' | '@' | '&' | '_' | '{' | '}' | '<' | '>' | '-' | ' ' | '	' | '`') +
   ;


STRING
   : '`' (~ ('`' | '\n'))* '`'
   ;


COMMENT
   : (';'|'#') ~ [\r\n]* EOL -> skip
   ;


EOL
   : [\r\n]
   ;


WS
   : [ \t] + -> skip
   ;
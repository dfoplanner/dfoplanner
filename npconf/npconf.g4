grammar npconf;

npconf
    : (section | EOL)*
    ;

section
   : sectionheader line*
   | sectionheader section* sectionfooter
   ;

sectionheader
   : '[' string ']' EOL
   ;

sectionfooter
   : '[/' string ']' EOL
   ;

string
   : CHARS
   | STRING
   ;

line
   : stringlist EOL
   ;

stringlist
   : string (',' string?)*
   ;


CHARS
   : ('A' .. 'Z' | '0' .. '9' | 'a' .. 'z' | '.' | '%' | '"' | '\\' | '/' | '*' | '@' | '&' | '_' | '{' | '}' | '<' | '>' | '-' | ' ' | '	') +
   ;


STRING
   : '`' (~ ('"' | '\n'))* '`'
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
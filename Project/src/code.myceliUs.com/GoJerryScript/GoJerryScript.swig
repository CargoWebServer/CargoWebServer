%module GoJerryScript
%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"
%}

%include "/usr/local/include/jerryscript-core.h"
%include "/usr/local/include/jerryscript-debugger.h"
%include "/usr/local/include/jerryscript-ext/handler.h"

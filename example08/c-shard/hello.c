// filename: _wale.c
#include "hello.h"
#include <stdio.h>

int main() {
  printf("This is a C Application.\n");
  GoString name = {"Jack", 4};
  SayHello(name);
  SayBye();
  return 0;
}

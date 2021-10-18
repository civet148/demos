#include "libhello.h"
#include <stdio.h>
#include <string.h>

 int main() {
   char *strCall = "clang call golang: "; //输入字符串
   int length = strlen(strCall); //获取输入字符串长度
   GoString strInput = {strCall, length};//go中的字符串类型在c中为GoString
   GoString strOutput; //libhello.so动态库hello()方法返回值
   strOutput = hello(strInput); //调用libhello.so动态库方法hello()
   if (strOutput.n <= 0) {
      printf("output string length %d invalid", strOutput.n);
      return -1;
   }
   int outLen = strOutput.n+1; //返回字符串长度+1
   char *outBuf = NULL; //返回字符串临时存储内存空间
   outBuf = malloc(outLen); //分配临时空间
   memset(outBuf, 0, outLen); //初始化内存
   strncpy(outBuf, strOutput.p, strOutput.n); //拷贝字符串到临时空间
   printf("golang return output: [%s]\n", outBuf); //打印返回字符串信息
   return 0;
}


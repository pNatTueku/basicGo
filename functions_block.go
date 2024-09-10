/*
// C block

#include <stdio.h?
int main(){
	printf("Start\n");
	{
	printf("Statement 1\n");
	printf("Statement 2\n");
	}
	printf("End");
	return 0;
}

*/

package main

import "fmt"

func main() {
	fmt.Println("Start")
	{
		fmt.Println("Statement 1")
		fmt.Println("Statement 2")
	}
	fmt.Println("End")
}

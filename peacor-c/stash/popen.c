#include <stdio.h>


int
main (void)
{
	FILE *fp;
	char *cmd;
	/* If you want to read output from cmd */
	fp = popen(cmd,"r"); 
	cmd = "cut -f 2- var2obs.tsv";
	char str[16];
		/* read output from cmd */
  while (!feof(fp))
	{
		fscanf(fp, "%s\n", str);   /* or other STDIO input functions */
		// outputs WITHOUT any call to print functions:
/*
sh: 1: 1I^HHPTI: not found
sh: 2: @: not found
-0.002069       1.305102        2.264105        -0.422982       1.855844
-4.547360       0.398718        2.397524        4.312785        2.438333
1.274580        1.812472        -1.675880       1.038524        2.550305
*/
	}
	fclose(fp);

	/* If you want to send input to cmd */
	fp = popen(cmd,"w"); 
		/* write to cmd */
		//fprintf(fp,....);   /* or other STDIO output functions */
	fclose(fp);
/*###########################################################################*/

  return 0;
}

#include <stdio.h>
#include <string.h>


int
main (void)
{
	FILE *fp;
	char *fn = "var2obs.tsv";
	fp = fopen (fn, "r");
	char *dlm = "\t";
	int lnmx = 1024;
	char line[lnmx];
	while (fgets (line, lnmx, fp))
	{
		char *field;
		field = strtok (line, dlm);
		while (field)
		{
			printf ("%s\n", field);
			field = strtok (NULL, dlm);
		}
	}
  return 0;
}
